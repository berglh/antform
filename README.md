
#antform
---
[![Build Status](https://travis-ci.org/berglh/antform.svg?branch=master)](https://travis-ci.org/berglh/antform)<br />

Despite it's name, **antform** is not a form markup language for our *[arthropodous friends](https://www.youtube.com/watch?v=ZGIZ-zUvotM)*, it is actually a command line tool written in *Go* that converts a *[terraform](https://github.com/hashicorp/terraform) state* file into an *[ansible](https://github.com/ansible/ansible)* inventory file.

The idea is to use *terraform* to deploy your infrastructure and then configure it with *ansible* using the inventory produced by **antform**. It has currently only been tested with the *triton terraform* provider, so your mileage may vary.

- [Usage](#usage)
  - [Arguments](#arguments)
  - [Switches](#switches)
- [Examples](#examples)


## Usage

#### Arguments
Flag | Example | Description
:---:|:----|:---
`-f` | `/path/to/terraform.tfstate` | Specify the path to the terraform.tfstate file, defaults to current directory `terraform.tfstate`.
`-t` | `group` | Group the Terraform machines by tag, these exist in the attributes object and prefixed by `tags.`. This example would map to the key `tags.group`.

#### Switches
Flag | Description
:---:|:----
`-h`| Displays the usage information.

**Note:** Currently **antform** outputs to stdout and assumess that your `tfstate` file is in the same structure as the test [terraform.tfstate](terraform.tfstate) file in this repo. This solves the problem in my specific use case, but adding support for other *terraform* providers wouldn't be too dificult, however I haven't bothered to check if the provider structs are the same as *triton* so give it a go.


## Examples

**antform** outputs the data in your `tfstate` file, sorting each primary ip address by name:

```
~$ antform
[es-data-1]
192.160.0.1
[es-data-2]
192.160.0.2
[es-data-3]
192.160.0.3
```

**antform** can then split the machines by tags that you define in the your `terraform` machine configuration:

```
~$ antform
[elasticsearch]
192.160.0.1
192.160.0.2
192.160.0.3
```
