---
title: SDKs
description: Setup Control Points using SDK libraries
slug: setup-control-points-using-sdks-libraries
keywords:
  - setup
  - flow
  - control
  - points
  - sdk
sidebar_position: 3
---

```mdx-code-block
import {apertureVersion} from '../../../../apertureVersion.js';
import DocCardList from '@theme/DocCardList';
```

For services to communicate with Aperture Agent, [Control Points][flow-control]
must be set to describe where the Flows are happening.

One way to achieve that is through [Istio/Envoy integration][istio].

A second way, one that allows for more customization, is to use Aperture SDK to
set feature or traffic Control Points within services.

We provide <a
href={`https://github.com/fluxninja/aperture/tree/${apertureVersion}/sdks/`}>Aperture
SDKs</a> for popular languages, such as :-

- [Golang][golang]
- [Java][java]
- [JavaScript][javascript]

Aperture SDK allows you to manually wrap any function call or code snippet
inside the Service code as a Feature Control Point. Every invocation of the
Feature is a Flow from the perspective of Aperture.

It also provides middlewares for popular frameworks, allowing you to easily set
traffic Control Points within your service.

<DocCardList />

[flow-control]: /concepts/integrations/flow-control/flow-control.md
[istio]: /get-started/integrations/flow-control/envoy/istio.md
[golang]: ./go/manual.md
[java]: ./java/manual.md
[javascript]: ./javascript/manual.md
