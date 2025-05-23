data "aws_vpc" "gemtech-vpc" {
  filter {
    name   = "vpc-id"
    values = ["vpc-07a258558f9cfc7ed"]
  }
}

data "aws_subnet" "private" {
  for_each = toset(local.private_subnets)

  id = each.value
}