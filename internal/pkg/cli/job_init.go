// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"errors"
	"fmt"

	"github.com/aws/copilot-cli/internal/pkg/aws/sessions"
	"github.com/aws/copilot-cli/internal/pkg/cli/group"
	"github.com/aws/copilot-cli/internal/pkg/config"
	"github.com/aws/copilot-cli/internal/pkg/deploy/cloudformation"
	"github.com/aws/copilot-cli/internal/pkg/initworkload"
	"github.com/aws/copilot-cli/internal/pkg/manifest"
	"github.com/aws/copilot-cli/internal/pkg/term/color"
	"github.com/aws/copilot-cli/internal/pkg/term/log"
	termprogress "github.com/aws/copilot-cli/internal/pkg/term/progress"
	"github.com/aws/copilot-cli/internal/pkg/term/prompt"
	"github.com/aws/copilot-cli/internal/pkg/term/selector"
	"github.com/aws/copilot-cli/internal/pkg/workspace"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	jobInitSchedulePrompt = "How would you like to " + color.Emphasize("schedule") + " this job?"
	jobInitScheduleHelp   = `How to determine this job's schedule. "Rate" lets you define the time between 
executions and is good for jobs which need to run frequently. "Fixed Schedule"
lets you use a predefined or custom cron schedule and is good for less-frequent 
jobs or those which require specific execution schedules.`
)

const (
	fmtAddJobToAppStart    = "Creating ECR repositories for job %s."
	fmtAddJobToAppFailed   = "Failed to create ECR repositories for job %s.\n"
	fmtAddJobToAppComplete = "Created ECR repositories for job %s.\n"
)

const (
	job = "job"
)

type initJobOpts struct {
	initWkldVars

	// Interfaces to interact with dependencies.
	fs     afero.Fs
	prompt prompter
	sel    initJobSelector

	// Outputs stored on successful actions.
	manifestPath string
}

func newInitJobOpts(vars initWkldVars) (*initJobOpts, error) {
	store, err := config.NewStore()
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to config store: %w", err)
	}

	ws, err := workspace.New()
	if err != nil {
		return nil, fmt.Errorf("workspace cannot be created: %w", err)
	}

	p := sessions.NewProvider()
	sess, err := p.Default()
	if err != nil {
		return nil, err
	}

	prompter := prompt.New()
	sel, err := selector.NewWorkspaceSelect(prompter, store, ws)
	if err != nil {
		return nil, err
	}

	initJob := initworkload.NewJobInitializer(
		store,
		ws,
		termprogress.NewSpinner(),
		cloudformation.New(sess),
	)
	return &initJobOpts{
		initWkldVars: vars,

		fs:     &afero.Afero{Fs: afero.NewOsFs()},
		init:   initJob,
		prompt: prompter,
		sel:    sel,
	}, nil
}

// Validate returns an error if the flag values passed by the user are invalid.
func (o *initJobOpts) Validate() error {
	if o.wkldType != "" {
		if err := validateJobType(o.wkldType); err != nil {
			return err
		}
	}
	if o.name != "" {
		if err := validateJobName(o.name); err != nil {
			return err
		}
	}
	if o.dockerfilePath != "" && o.image != "" {
		return fmt.Errorf("--%s and --%s cannot be specified together", dockerFileFlag, imageFlag)
	}
	if o.dockerfilePath != "" {
		if _, err := o.fs.Stat(o.dockerfilePath); err != nil {
			return err
		}
	}
	if o.schedule != "" {
		if err := validateSchedule(o.schedule); err != nil {
			return err
		}
	}
	if o.timeout != "" {
		if err := validateTimeout(o.timeout); err != nil {
			return err
		}
	}
	if o.retries < 0 {
		return errors.New("number of retries must be non-negative")
	}
	return nil
}

// Ask prompts for fields that are required but not passed in.
func (o *initJobOpts) Ask() error {
	if err := o.askJobType(); err != nil {
		return err
	}
	if err := o.askJobName(); err != nil {
		return err
	}
	dfSelected, err := o.askDockerfile()
	if err != nil {
		return err
	}
	if !dfSelected {
		if err := o.askImage(); err != nil {
			return err
		}
	}
	if err := o.askSchedule(); err != nil {
		return err
	}
	return nil
}

// Execute writes the job's manifest file, creates an ECR repo, and stores the name in SSM.
func (o *initJobOpts) Execute() error {
	manifestPath, err := o.init.Job(&initworkload.WorkloadProps{
		App:            o.appName,
		Name:           o.name,
		DockerfilePath: o.dockerfilePath,
		Image:          o.image,
		Type: o.type,
		
		Schedule: o.schedule,
		Retries: o.retries,
		Timeout: o.timeout,
	})
	if err != nil {
		return err
	}
	o.manifestPath = manifestPath

	return nil
}

func (o *initJobOpts) askJobType() error {
	if o.wkldType != "" {
		return nil
	}
	// short circuit since there's only one valid job type.
	o.wkldType = manifest.ScheduledJobType
	return nil
}

func (o *initJobOpts) askJobName() error {
	if o.name != "" {
		return nil
	}

	name, err := o.prompt.Get(
		fmt.Sprintf(fmtWkldInitNamePrompt, color.Emphasize("name"), color.HighlightUserInput(o.wkldType)),
		fmt.Sprintf(fmtWkldInitNameHelpPrompt, job, o.appName),
		validateSvcName,
		prompt.WithFinalMessage("Job name:"),
	)
	if err != nil {
		return fmt.Errorf("get job name: %w", err)
	}
	o.name = name
	return nil
}

func (o *initJobOpts) askImage() error {
	if o.image != "" {
		return nil
	}
	image, err := o.prompt.Get(wkldInitImagePrompt, wkldInitImagePromptHelp, nil,
		prompt.WithFinalMessage("Image:"))
	if err != nil {
		return fmt.Errorf("get image location: %w", err)
	}
	o.image = image
	return nil
}

// isDfSelected indicates if any Dockerfile is in use.
func (o *initJobOpts) askDockerfile() (isDfSelected bool, err error) {
	if o.dockerfilePath != "" || o.image != "" {
		return true, nil
	}
	df, err := o.sel.Dockerfile(
		fmt.Sprintf(fmtWkldInitDockerfilePrompt, color.HighlightUserInput(o.name)),
		fmt.Sprintf(fmtWkldInitDockerfilePathPrompt, color.HighlightUserInput(o.name)),
		wkldInitDockerfileHelpPrompt,
		wkldInitDockerfilePathHelpPrompt,
		func(v interface{}) error {
			return validatePath(afero.NewOsFs(), v)
		},
	)
	if err != nil {
		return false, fmt.Errorf("select Dockerfile: %w", err)
	}
	if df == selector.DockerfilePromptUseImage {
		return false, nil
	}
	o.dockerfilePath = df
	return true, nil
}

func (o *initJobOpts) askSchedule() error {
	if o.schedule != "" {
		return nil
	}
	schedule, err := o.sel.Schedule(
		jobInitSchedulePrompt,
		jobInitScheduleHelp,
		validateSchedule,
		validateRate,
	)
	if err != nil {
		return fmt.Errorf("get schedule: %w", err)
	}

	o.schedule = schedule
	return nil
}

// RecommendedActions returns follow-up actions the user can take after successfully executing the command.
func (o *initJobOpts) RecommendedActions() []string {
	return []string{
		fmt.Sprintf("Update your manifest %s to change the defaults.", color.HighlightResource(o.manifestPath)),
		fmt.Sprintf("Run %s to deploy your job to a %s environment.",
			color.HighlightCode(fmt.Sprintf("copilot job deploy --name %s --env %s", o.name, defaultEnvironmentName)),
			defaultEnvironmentName),
	}
}

// buildJobInitCmd builds the command for creating a new job.
func buildJobInitCmd() *cobra.Command {
	vars := initWkldVars{}
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Creates a new scheduled job in an application.",
		Example: `
  Create a "reaper" scheduled task to run once per day.
  /code $ copilot job init --name reaper --dockerfile ./frontend/Dockerfile --schedule "every 2 hours"

  Create a "report-generator" scheduled task with retries.
  /code $ copilot job init --name report-generator --schedule "@monthly" --retries 3 --timeout 900s`,
		RunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			opts, err := newInitJobOpts(vars)
			if err != nil {
				return err
			}
			if err := opts.Validate(); err != nil { // validate flags
				return err
			}
			if err := opts.Ask(); err != nil {
				return err
			}
			if err := opts.Execute(); err != nil {
				return err
			}
			log.Infoln("Recommended follow-up actions:")
			for _, followup := range opts.RecommendedActions() {
				log.Infof("- %s\n", followup)
			}
			return nil
		}),
	}
	cmd.Flags().StringVarP(&vars.appName, appFlag, appFlagShort, tryReadingAppName(), appFlagDescription)
	cmd.Flags().StringVarP(&vars.name, nameFlag, nameFlagShort, "", jobFlagDescription)
	cmd.Flags().StringVarP(&vars.wkldType, jobTypeFlag, jobTypeFlagShort, "", jobTypeFlagDescription)
	cmd.Flags().StringVarP(&vars.dockerfilePath, dockerFileFlag, dockerFileFlagShort, "", dockerFileFlagDescription)
	cmd.Flags().StringVarP(&vars.schedule, scheduleFlag, scheduleFlagShort, "", scheduleFlagDescription)
	cmd.Flags().StringVar(&vars.timeout, timeoutFlag, "", timeoutFlagDescription)
	cmd.Flags().IntVar(&vars.retries, retriesFlag, 0, retriesFlagDescription)
	cmd.Flags().StringVarP(&vars.image, imageFlag, imageFlagShort, "", imageFlagDescription)

	cmd.Annotations = map[string]string{
		"group": group.Develop,
	}

	return cmd
}
