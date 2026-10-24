package featureflag

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures Feature Flags resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_feature_flag", func(r *config.Resource) {
		r.Kind = "FeatureFlag"
		r.ShortGroup = "featureflag"
	})
	p.AddResourceConfigurator("datadog_feature_flag_environment", func(r *config.Resource) {
		r.Kind = "FeatureFlagEnvironment"
		r.ShortGroup = "featureflag"
	})
	p.AddResourceConfigurator("datadog_feature_flag_environment_config", func(r *config.Resource) {
		r.Kind = "FeatureFlagEnvironmentConfig"
		r.ShortGroup = "featureflag"
		r.References["feature_flag_id"] = config.Reference{TerraformName: "datadog_feature_flag"}
		r.References["environment_id"] = config.Reference{TerraformName: "datadog_feature_flag_environment"}
	})
	p.AddResourceConfigurator("datadog_feature_flag_allocations", func(r *config.Resource) {
		r.Kind = "FeatureFlagAllocations"
		r.ShortGroup = "featureflag"
		r.References["feature_flag_id"] = config.Reference{TerraformName: "datadog_feature_flag"}
		r.References["environment_id"] = config.Reference{TerraformName: "datadog_feature_flag_environment"}
	})
}
