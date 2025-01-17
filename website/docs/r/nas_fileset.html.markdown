---
subcategory: "Network Attached Storage (NAS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nas_fileset"
sidebar_current: "docs-alicloud-resource-nas-fileset"
description: |-
  Provides a Alicloud Network Attached Storage (NAS) Fileset resource.
---

# alicloud\_nas\_fileset

Provides a Network Attached Storage (NAS) Fileset resource.

For information about Network Attached Storage (NAS) Fileset and how to use it, see [What is Fileset](https://www.alibabacloud.com/help/en/doc-detail/27530.html).

-> **NOTE:** Available in v1.153.0+.

## Example Usage

Basic Usage

```terraform
data "alicloud_nas_zones" "example" {
  file_system_type = "cpfs"
}

resource "alicloud_vpc" "example" {
  vpc_name   = "terraform-example"
  cidr_block = "172.17.3.0/24"
}

resource "alicloud_vswitch" "example" {
  vswitch_name = "terraform-example"
  cidr_block   = "172.17.3.0/24"
  vpc_id       = alicloud_vpc.example.id
  zone_id      = data.alicloud_nas_zones.example.zones[1].zone_id
}

resource "alicloud_nas_file_system" "example" {
  protocol_type    = "cpfs"
  storage_type     = "advance_200"
  file_system_type = "cpfs"
  capacity         = 3600
  description      = "terraform-example"
  zone_id          = data.alicloud_nas_zones.example.zones[1].zone_id
  vpc_id           = alicloud_vpc.example.id
  vswitch_id       = alicloud_vswitch.example.id
}

resource "alicloud_nas_fileset" "example" {
  file_system_id   = alicloud_nas_file_system.example.id
  description      = "terraform-example"
  file_system_path = "/example_path/"
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) The description of the Fileset. It must be `2` to `128` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
* `dry_run` - (Optional) The dry run.
* `file_system_id` - (Required, ForceNew) The ID of the file system.
* `file_system_path` - (Required, ForceNew) The path of the fileset.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID of the file set. The value formats as `<file_system_id>:<fileset_id>`.
* `fileset_id` - The first ID of the resource.
* `status` - The status of the fileset. 

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 1 mins) Used when create the fileset.
* `delete` - (Defaults to 1 mins) Used when delete the fileset.

## Import

Network Attached Storage (NAS) Fileset can be imported using the id, e.g.

```shell
$ terraform import alicloud_nas_fileset.example <file_system_id>:<fileset_id>
```