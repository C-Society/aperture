---
title: Policy
sidebar_position: 1
---

:::info

See also [Policy reference](/reference/policies/spec.md#policy)

:::

The Policy language enables expression of [Control Systems][control-system] as
code. This language can be used by service operators to express reliability
automation workflows for each service.

Policy specification consists of two parts

1. A [Circuit][circuit] that expresses the Control System as an execution graph.
2. A list of [Resources][resources] which need to be set up in order to support
   the Circuit.

Aperture comes with a pre-packaged [blueprints][blueprints] that can be used
both as a guide for creating new policies, or used as-is.

Policies provide a framework for defining and managing reliability criteria, and
conditions as code. It's a way of enforcing reliability policies
programmatically, running in a continuous control loop. In an application
reliability context, it codifies the capability of the application to modify its
operational state to achieve the best possible mode of operation despite
overload and failures.

[circuit]: /concepts/policy/circuit.md
[resources]: /concepts/policy/resources.md
[blueprints]: /reference/policies/bundled-blueprints/bundled-blueprints.md
[control-system]: https://en.wikipedia.org/wiki/Control_system
