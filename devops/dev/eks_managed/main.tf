provider "aws" {
  region = local.region
}

terraform {
  backend "s3" {
    bucket  = "gemtech-remotestate-dev"
    key     = "c2-eks/eks-managed/terraform.tfstate"
    region  = "us-east-1"
    encrypt = true
    profile = "default"
  }
}

locals {
  name                  = "c2-eks"
  region                = "us-east-1"
  cidr_block        = data.aws_vpc.gemtech-vpc.cidr_block
  additional_cidr_block = "172.16.0.0/16"
  environment           = "test"
  label_order           = ["name", "environment"]
  tags = {
    "kubernetes.io/cluster/${module.eks.cluster_name}" = "owned"
  }
  private_subnets     = ["subnet-0093f13ff04f1ecab", "subnet-0e629b4e20cd1f521"]
}


# ################################################################################
# Security Groups module call
################################################################################

module "ssh" {
  source = "git::ssh://git@github.com/mauricetjmurphy/gemtech-terraform-modules.git//security-group"
  name        = "${local.name}-ssh"
  environment = local.environment
  vpc_id      = data.aws_vpc.gemtech-vpc.id
  new_sg_ingress_rules_with_cidr_blocks = [{
    rule_count  = 1
    from_port   = 22
    protocol    = "tcp"
    to_port     = 22
    cidr_blocks = [local.cidr_block, local.additional_cidr_block]
    description = "Allow ssh traffic."
    },
    {
      rule_count  = 2
      from_port   = 27017
      protocol    = "tcp"
      to_port     = 27017
      cidr_blocks = [local.additional_cidr_block]
      description = "Allow Mongodb traffic."
    }
  ]

  ## EGRESS Rules
  new_sg_egress_rules_with_cidr_blocks = [{
    rule_count  = 1
    from_port   = 22
    protocol    = "tcp"
    to_port     = 22
    cidr_blocks = [local.cidr_block, local.additional_cidr_block]
    description = "Allow ssh outbound traffic."
    },
    {
      rule_count  = 2
      from_port   = 27017
      protocol    = "tcp"
      to_port     = 27017
      cidr_blocks = [local.additional_cidr_block]
      description = "Allow Mongodb outbound traffic."
  }]
}

module "http_https" {
  source = "git::ssh://git@github.com/mauricetjmurphy/gemtech-terraform-modules.git//security-group"
  name        = "${local.name}-http-https"
  environment = local.environment

  vpc_id = data.aws_vpc.gemtech-vpc.id
  ## INGRESS Rules
  new_sg_ingress_rules_with_cidr_blocks = [{
    rule_count  = 1
    from_port   = 22
    protocol    = "tcp"
    to_port     = 22
    cidr_blocks = [local.cidr_block]
    description = "Allow ssh traffic."
    },
    {
      rule_count  = 2
      from_port   = 80
      protocol    = "tcp"
      to_port     = 80
      cidr_blocks = [local.cidr_block]
      description = "Allow http traffic."
    },
    {
      rule_count  = 3
      from_port   = 443
      protocol    = "tcp"
      to_port     = 443
      cidr_blocks = [local.cidr_block]
      description = "Allow https traffic."
    }
  ]

  ## EGRESS Rules
  new_sg_egress_rules_with_cidr_blocks = [{
    rule_count       = 1
    from_port        = 0
    protocol         = "-1"
    to_port          = 0
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
    description      = "Allow all traffic."
    }
  ]
}

################################################################################
# KMS Module call
################################################################################
# module "kms" {
#   source = "git::ssh://git@github.com/mauricetjmurphy/gemtech-terraform-modules.git//security-kms"
#   name                = "${local.name}-kms"
#   environment         = local.environment
#   label_order         = local.label_order
#   enabled             = true
#   description         = "KMS key for EBS of EKS nodes"
#   enable_key_rotation = false
#   policy              = data.aws_iam_policy_document.kms.json
# }

# data "aws_iam_policy_document" "kms" {
#   version = "2012-10-17"
#   statement {
#     sid    = "Enable IAM User Permissions"
#     effect = "Allow"
#     principals {
#       type        = "AWS"
#       identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
#     }
#     actions   = ["kms:*"]
#     resources = ["*"]
#   }
# }

################################################################################
# Customer-Managed KMS Key
################################################################################
# resource "aws_kms_key" "my_key" {
#   description         = "Customer managed KMS key for EBS encryption"
#   enable_key_rotation = true
#   policy              = data.aws_iam_policy_document.kms.json

#   tags = {
#     Name        = "${local.name}-kms"
#     Environment = local.environment
#   }
# }

# data "aws_iam_policy_document" "kms" {
#   version = "2012-10-17"

#   statement {
#     sid    = "Enable IAM User Permissions"
#     effect = "Allow"
#     principals {
#       type        = "AWS"
#       identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
#     }
#     actions   = ["kms:*"]
#     resources = ["*"]
#   }

#   statement {
#     sid    = "Allow use of the key"
#     effect = "Allow"
#     actions = [
#       "kms:Encrypt",
#       "kms:Decrypt",
#       "kms:ReEncrypt*",
#       "kms:GenerateDataKey*",
#       "kms:DescribeKey"
#     ]
#     resources = ["*"]
#   }

#   statement {
#     sid    = "Allow attachment of persistent resources"
#     effect = "Allow"
#     actions = [
#       "kms:CreateGrant",
#       "kms:ListGrants",
#       "kms:RevokeGrant"
#     ]
#     resources = ["*"]
#     condition {
#       test     = "Bool"
#       variable = "kms:GrantIsForAWSResource"
#       values   = ["true"]
#     }
#   }
# }

data "aws_caller_identity" "current" {}

################################################################################
# EKS Module call 
################################################################################
module "eks" {
  source  = "git::ssh://git@github.com/mauricetjmurphy/gemtech-terraform-modules.git//eks"
  enabled = true

  name        = local.name
  environment = local.environment
  label_order = local.label_order

  # EKS
  kubernetes_version     = "1.28"
  endpoint_public_access = true
  # Networking
  vpc_id                            = data.aws_vpc.gemtech-vpc.id
  subnet_ids                        = local.private_subnets
  allowed_security_groups           = [module.ssh.security_group_id]
  eks_additional_security_group_ids = ["${module.ssh.security_group_id}", "${module.http_https.security_group_id}"]
  allowed_cidr_blocks               = [local.cidr_block]

  # AWS Managed Node Group
  # Node Groups Defaults Values It will Work all Node Groups
  managed_node_group_defaults = {
    subnet_ids                          = local.private_subnets
    nodes_additional_security_group_ids = [module.ssh.security_group_id]
    tags = {
      "kubernetes.io/cluster/${module.eks.cluster_name}" = "shared"
      "k8s.io/cluster/${module.eks.cluster_name}"        = "shared"
    }
    block_device_mappings = {
      xvda = {
        device_name = "/dev/xvda"
        ebs = {
          volume_size = 50
          volume_type = "gp3"
          iops        = 3000
          throughput  = 150
          encrypted   = false
          # kms_key_id  = aws_kms_key.my_key.arn
        }
      }
    }
  }
  managed_node_group = {
    service-01 = {
      name           = "${module.eks.cluster_name}-service-01"
      capacity_type  = "ON_DEMAND"
      min_size       = 2
      max_size       = 4
      desired_size   = 2
      instance_types = ["t3.small"]
    }

    service-02 = {
      name                 = "${module.eks.cluster_name}-service-02"
      capacity_type        = "ON_DEMAND"
      min_size             = 2
      max_size             = 4
      desired_size         = 2
      force_update_version = true
      instance_types       = ["t3.small"]
    }
  }

  apply_config_map_aws_auth = true
  map_additional_iam_users = [
    {
      userarn  = "arn:aws:iam::144817152095:user/mauricetjmurphy"
      username = "mauricetjmurphy@gmail.com"
      groups   = ["system:masters"]
    }
  ]
}

## Kubernetes provider configuration
data "aws_eks_cluster" "this" {
  depends_on = [module.eks]
  name       = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "this" {
  depends_on = [module.eks]
  name       = module.eks.cluster_certificate_authority_data
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.this.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.this.certificate_authority[0].data)
  token                  = data.aws_eks_cluster_auth.this.token
}

# Use kubernetes_manifest to apply YAML files, depending on node labeling
resource "kubernetes_manifest" "main" {
  for_each = fileset("k8s", "*.yaml")

  manifest = yamldecode(file("k8s/${each.value}"))
}