package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAsciiFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::ascii("Hello 世界!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "Hello !"),
				),
			},
			{
				Config: `
				output "test" {
					value = provider::curious::ascii("Café résumé")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "Cafe resume"),
				),
			},
		},
	})
}

func TestFlatFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::flat("Hello World!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "helloworld"),
				),
			},
		},
	})
}

func TestKebabFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::kebab("Hello World!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "hello-world"),
				),
			},
		},
	})
}

func TestCamelFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::camel("Hello World!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "helloWorld"),
				),
			},
		},
	})
}

func TestPascalFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::pascal("Hello World!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "HelloWorld"),
				),
			},
		},
	})
}

func TestSnakeFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::snake("Hello World!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "hello_world"),
				),
			},
		},
	})
}

func TestUpperFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::upper("Hello World!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "HELLO_WORLD"),
				),
			},
		},
	})
}

func TestTrainFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::train("Hello World!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "HELLO-WORLD"),
				),
			},
		},
	})
}

func TestAdaFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::ada("Hello World!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "Hello_World"),
				),
			},
		},
	})
}

func TestEliteFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::elite("Sponge Bob!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "SPoNGe BoB!"),
				),
			},
			{
				Config: `
				output "test" {
					value = provider::curious::elite("Café")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "CaFé"),
				),
			},
		},
	})
}

func TestSpongeFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::sponge("sponge bob")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "sPoNgE bOb"),
				),
			},
			{
				Config: `
				output "test" {
					value = provider::curious::sponge("Café-World")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "cAfÉ-wOrLd"),
				),
			},
		},
	})
}
func TestAsciiPrintableFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::ascii_printable("Hello\t\nWorld!")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "HelloWorld!"),
				),
			},
			{
				Config: `
				output "test" {
					value = provider::curious::ascii_printable("Café 世界! Test")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "Cafe ! Test"),
				),
			},
		},
	})
}

func TestLatinizeFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::curious::latinize("räksmörgås")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "raksmorgas"),
				),
			},
		},
	})
}
