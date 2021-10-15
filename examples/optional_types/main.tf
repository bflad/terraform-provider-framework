terraform {
  required_providers {
    framework = {
      source = "bflad/framework"
    }
  }
  required_version = "1.0.9"
}

resource "framework_optional_types" "example" {}
