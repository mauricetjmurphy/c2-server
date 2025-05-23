module "load_balancer_controller" {
  source = "git::ssh://git@github.com/mauricetjmurphy/gemtech-terraform-modules.git//eks-lb-controller"
  enabled = true
  cluster_identity_oidc_issuer     = data.aws_eks_cluster.cluster.cluster_oidc_issuer_url
  cluster_identity_oidc_issuer_arn = data.aws_eks_cluster.cluster.oidc_provider_arn
  cluster_name                     = data.aws_eks_cluster.cluster.cluster_id

  depends_on = [data.aws_eks_cluster.cluster]
}