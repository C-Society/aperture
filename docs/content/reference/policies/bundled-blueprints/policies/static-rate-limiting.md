---
title: Static Rate Limiting Policy
---

```mdx-code-block
import {apertureVersion} from '../../../../apertureVersion.js';
```

## Introduction

This blueprint provides a simple static rate limiting policy and a dashboard.

## Links

Sample values file: <a
href={`https://github.com/fluxninja/aperture/tree/${apertureVersion}/blueprints/policies/static-rate-limiting/values.yaml`}>values.yaml</a>

Sample values file (required fields only): <a
href={`https://github.com/fluxninja/aperture/tree/${apertureVersion}/blueprints/policies/static-rate-limiting/values_required.yaml`}>values_required.yaml</a>

Code: <a
href={`https://github.com/fluxninja/aperture/tree/${apertureVersion}/blueprints/policies/static-rate-limiting`}>static-rate-limiting</a>

## Configuration

<!-- Configuration Marker -->

```mdx-code-block
export const ParameterHeading = ({children}) => (
  <span style={{fontWeight: "bold"}}>{children}</span>
);

export const WrappedDescription = ({children}) => (
  <span style={{wordWrap: "normal"}}>{children}</span>
);

export const RefType = ({type, reference}) => (
  <a href={reference}>{type}</a>
);

export const ParameterDescription = ({name, type, reference, value, description}) => (
  <table class="blueprints-params">
  <tr>
    <td><ParameterHeading>Parameter</ParameterHeading></td>
    <td><code>{name}</code></td>
  </tr>
  <tr>
    <td><ParameterHeading>Type</ParameterHeading></td>
    <td><em>{reference == "" ? type : <RefType type={type} reference={reference} />}</em></td>
  </tr>
  <tr>
    <td class="blueprints-default-heading"><ParameterHeading>Default Value</ParameterHeading></td>
    <td><code>{value}</code></td>
  </tr>
  <tr>
    <td class="blueprints-description"><ParameterHeading>Description</ParameterHeading></td>
    <td class="blueprints-description"><WrappedDescription>{description}</WrappedDescription></td>
  </tr>
</table>
);
```

<h3 class="blueprints-h3">Common</h3>

<ParameterDescription
    name="common.policy_name"
    type="string"
    reference=""
    value="__REQUIRED_FIELD__"
    description='Name of the policy.' />

<h3 class="blueprints-h3">Policy</h3>

<ParameterDescription
    name="policy.evaluation_interval"
    type="string"
    reference=""
    value="'300s'"
    description='How often should the policy be re-evaluated' />

<ParameterDescription
    name="policy.classifiers"
    type="[]aperture.spec.v1.Classifier"
    reference="../../spec#v1-classifier"
    value="[]"
    description='List of classification rules.' />

<h4 class="blueprints-h4">Rate Limiter</h4>

<ParameterDescription
    name="policy.rate_limiter.rate_limit"
    type="float64"
    reference=""
    value="__REQUIRED_FIELD__"
    description='Number of requests per `policy.rate_limiter.parameters.limit_reset_interval` to accept' />

<ParameterDescription
    name="policy.rate_limiter.flow_selector"
    type="aperture.spec.v1.FlowSelector"
    reference="../../spec#v1-flow-selector"
    value="{'flow_matcher': {'control_point': '__REQUIRED_FIELD__'}, 'service_selector': {'agent_group': 'default', 'service': '__REQUIRED_FIELD__'}}"
    description='A flow selector to match requests against' />

<ParameterDescription
    name="policy.rate_limiter.flow_selector.service_selector.service"
    type="string"
    reference=""
    value="__REQUIRED_FIELD__"
    description='Service Name.' />

<ParameterDescription
    name="policy.rate_limiter.flow_selector.flow_matcher.control_point"
    type="string"
    reference=""
    value="__REQUIRED_FIELD__"
    description='Control Point Name.' />

<ParameterDescription
    name="policy.rate_limiter.parameters"
    type="aperture.spec.v1.RateLimiterParameters"
    reference="../../spec#v1-rate-limiter-parameters"
    value="{'label_key': '__REQUIRED_FIELD__', 'lazy_sync': {'enabled': True, 'num_sync': 5}, 'limit_reset_interval': '1s'}"
    description='Parameters.' />

<ParameterDescription
    name="policy.rate_limiter.parameters.label_key"
    type="string"
    reference=""
    value="__REQUIRED_FIELD__"
    description='Flow label to use for rate limiting.' />

<ParameterDescription
    name="policy.rate_limiter.dynamic_config"
    type="aperture.spec.v1.RateLimiterDefaultConfig"
    reference="../../spec#v1-rate-limiter-default-config"
    value="{'overrides': []}"
    description='Dynamic configuration for rate limiter that can be applied at the runtime.' />

<h3 class="blueprints-h3">Dashboard</h3>

<ParameterDescription
    name="dashboard.refresh_interval"
    type="string"
    reference=""
    value="'10s'"
    description='Refresh interval for dashboard panels.' />

<h4 class="blueprints-h4">Datasource</h4>

<ParameterDescription
    name="dashboard.datasource.name"
    type="string"
    reference=""
    value="'$datasource'"
    description='Datasource name.' />

<ParameterDescription
    name="dashboard.datasource.filter_regex"
    type="string"
    reference=""
    value="''"
    description='Datasource filter regex.' />