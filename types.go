package main

// Metadata is generic metadata applicable to a wide variety of types
type Metadata struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

// Workspace holds information about a workspace
type Workspace struct {
	Metadata
	ProjectIds []string                `json:"project_ids,omitempty"`
	Inputs     map[string]ProjectInput `json:"inputs,omitempty"`
}

// Project holds information about a project
type Project struct {
	Metadata
	Inputs  map[string]ProjectInput  `json:"inputs,omitempty"`
	Outputs map[string]ProjectOutput `json:"outputs,omitempty"`
	Status  map[ProjectStatus]bool   `json:"status,omitempty"`
}

// ProjectInput holds information about a single input of a project
type ProjectInput struct {
	Metadata
	Type           ProjectInputType `json:"type"`
	NormalizedName string           `json:"normalized_name"`
}

// ProjectOutput holds information about a single output of a project
type ProjectOutput struct {
	Metadata
	Status ProjectOutputStatus `json:"status"`
}

// ProjectStatus stores the current project status
type ProjectStatus string

const (
	// ProjectStatusInputSources indicates the project has source folder uploaded
	ProjectStatusInputSources ProjectStatus = "sources"
	// ProjectStatusInputCustomizations indicates the project has customizations folder uploaded
	ProjectStatusInputCustomizations ProjectStatus = "customizations"
	// ProjectStatusInputConfigs indicates the project has configs
	ProjectStatusInputConfigs ProjectStatus = "configs"
	// ProjectStatusInputReference indicates the project has references to workspace level inputs
	ProjectStatusInputReference ProjectStatus = "reference"
	// ProjectStatusPlanning indicates the project is currently generating a plan
	ProjectStatusPlanning ProjectStatus = "planning"
	// ProjectStatusPlan indicates the project has a plan
	ProjectStatusPlan ProjectStatus = "plan"
	// ProjectStatusStalePlan indicates that the inputs have changed after the plan was last generated
	ProjectStatusStalePlan ProjectStatus = "stale_plan"
	// ProjectStatusPlanError indicates that an error occurred during planning
	ProjectStatusPlanError ProjectStatus = "plan_error"
	// ProjectStatusOutputs indicates the project has project artifacts generated
	ProjectStatusOutputs ProjectStatus = "outputs"
)

// ProjectOutputStatus is the status of a project output
type ProjectOutputStatus string

const (
	// ProjectOutputStatusInProgress indicates that the transformation is ongoing
	ProjectOutputStatusInProgress = "transforming"
	// ProjectOutputStatusDoneSuccess indicates that the transformation completed successfully
	ProjectOutputStatusDoneSuccess = "done"
	// ProjectOutputStatusDoneError indicates an error like if the transformation was cancelled or the timeout expired
	ProjectOutputStatusDoneError = "error"
)

// ProjectInputType is the type of the project input
type ProjectInputType string

const (
	// ProjectInputSources is the type for project inputs that are folders containing source code
	ProjectInputSources ProjectInputType = ProjectInputType(ProjectStatusInputSources)
	// ProjectInputCustomizations is the type for project inputs that are folders containing customization files
	ProjectInputCustomizations ProjectInputType = ProjectInputType(ProjectStatusInputCustomizations)
	// ProjectInputConfigs is the type for project inputs that are config files
	ProjectInputConfigs ProjectInputType = ProjectInputType(ProjectStatusInputConfigs)
	// ProjectInputReference is the type for project inputs that are references to workspace level inputs
	ProjectInputReference ProjectInputType = ProjectInputType(ProjectStatusInputReference)
)

const (
	// WORKSPACES_BUCKET is the workspaces bucket
	WORKSPACES_BUCKET = "workspaces"
	// PROJECTS_BUCKET is the projects bucket
	PROJECTS_BUCKET = "projects"
)
