package v1alpha1

// JobsMetrics return a set of ci metric for every repository in quality studio database
type JobsMetrics struct {
	// Indicate a GitHub repository name. "e2e-tests", "kubernetes"
	RepositoryName string `json:"repository_name"`

	// PostSubmits, periodics or presubmits
	JobType string `json:"type"`

	// Indicate a GitHub organization name
	GitOrganization string `json:"git_organization"`

	// A set of Jobs
	Jobs []Jobs `json:"jobs"`
}

// Metrics for specific job name
type Jobs struct {
	// Name of a prow job
	Name string `json:"name"`

	// Metadata about jobs
	Summary Summary `json:"summary"`

	// Set of metrics about job
	Metrics []Metrics `json:"metrics"`
}

// Set of Metrics for a specific job
type Metrics struct {
	// Return the count of successful jobs
	SuccessCount float64 `json:"success_count"`

	// Return the count of failed jobs
	FailureCount float64 `json:"failure_count"`

	// Return the count of failed jobs due to e2e tests errors
	JobFailedByE2ETestsCount float64 `json:"failure_by_e2e_tests_count"`

	// Return the count of failed jobs due to build errors like installation errors
	JobFailedByBuildErrorsCount float64 `json:"failure_by_build_errors_count"`

	// Return the count of jobs that were not scheduled
	JobNotScheduledCount float64 `json:"not_scheduled_count"`

	// Return the total number of jobs
	TotalJobs float64 `json:"total_jobs"`

	// Return the date
	Date string `json:"date"`
}

type Summary struct {
	DateFrom string `json:"date_from"`

	DateTo string `json:"date_to"`

	SuccessCount float64 `json:"success_count"`

	JobFailedCount float64 `json:"failure_count"`

	JobFailedByE2ETestsCount float64 `json:"failure_by_e2e_tests_count"`

	JobFailedByBuildErrorsCount float64 `json:"failure_by_build_errors_count"`

	JobNotScheduledCount float64 `json:"not_scheduled_count"`

	TotalJobs float64 `json:"total_jobs"`
}
