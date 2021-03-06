package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAwsRouteTables(t *testing.T) {
	rInt := acctest.RandIntRange(0, 256)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsRouteTablesConfig(rInt),
			},
			{
				Config: testAccDataSourceAwsRouteTablesConfigWithDataSource(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.aws_route_tables.test", "ids.#", "4"),
					resource.TestCheckResourceAttr("data.aws_route_tables.private", "ids.#", "2"),
					resource.TestCheckResourceAttr("data.aws_route_tables.test2", "ids.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceAwsRouteTablesConfigWithDataSource(rInt int) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "172.%d.0.0/16"

  tags {
    Name = "terraform-testacc-route-tables-data-source"
  }
}

resource "aws_vpc" "test2" {
  cidr_block = "172.%d.0.0/16"

  tags {
    Name = "terraform-test2acc-route-tables-data-source"
  }
}

resource "aws_route_table" "test_public_a" {
  vpc_id            = "${aws_vpc.test.id}"

  tags {
    Name = "tf-acc-route-tables-data-source-public-a"
    Tier = "Public"
  }
}

resource "aws_route_table" "test_private_a" {
  vpc_id            = "${aws_vpc.test.id}"

  tags {
    Name = "tf-acc-route-tables-data-source-private-a"
    Tier = "Private"
  }
}

resource "aws_route_table" "test_private_b" {
  vpc_id            = "${aws_vpc.test.id}"

  tags {
    Name = "tf-acc-route-tables-data-source-private-b"
    Tier = "Private"
  }
}

data "aws_route_tables" "test" {
  vpc_id = "${aws_vpc.test.id}"
}

data "aws_route_tables" "test2" {
  vpc_id = "${aws_vpc.test2.id}"
}

data "aws_route_tables" "private" {
  vpc_id = "${aws_vpc.test.id}"
  tags {
    Tier = "Private"
  }
}
`, rInt, rInt)
}

func testAccDataSourceAwsRouteTablesConfig(rInt int) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "172.%d.0.0/16"

  tags {
    Name = "terraform-testacc-route-tables-data-source"
  }
}

resource "aws_route_table" "test_public_a" {
  vpc_id            = "${aws_vpc.test.id}"

  tags {
    Name = "tf-acc-route-tables-data-source-public-a"
    Tier = "Public"
  }
}

resource "aws_route_table" "test_private_a" {
  vpc_id            = "${aws_vpc.test.id}"

  tags {
    Name = "tf-acc-route-tables-data-source-private-a"
    Tier = "Private"
  }
}

resource "aws_route_table" "test_private_b" {
  vpc_id            = "${aws_vpc.test.id}"

  tags {
    Name = "tf-acc-route-tables-data-source-private-b"
    Tier = "Private"
  }
}
`, rInt)
}
