---
page_title: example Function - terraform-provider-framework
description: |-
  Echoes given argument as result.
---

# Function: example

Echoes given argument as result.

## Example Usage

```terraform
# result: testvalue
provider::framework::example("testvalue")
```

## Signature

```text
example(input string) string
```

## Arguments

1. `input` (String) String to echo
