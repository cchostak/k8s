data "aws_eks_cluster" "cluster" {
  name = module.my-cluster.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.my-cluster.cluster_id
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
  token                  = data.aws_eks_cluster_auth.cluster.token
  load_config_file       = false
  version                = "~> 1.9"
}

module "my-cluster" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = "my-cluster"
  cluster_version = "1.14"
  subnets         = ["subnet-abcde012", "subnet-bcde012a", "subnet-fghi345a"]
  vpc_id          = "vpc-1234556abcdef"

  worker_groups = [
    {
      instance_type = "m4.large"
      asg_max_size  = 5
    }
  ]
}

module "vpc" {
  source = "nekochans/vpc/aws"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  nat_instance_ami         = "ami-0af1df87db7b650f4"
  nat_instance_type        = "t2.micro"
  nat_instance_volume_type = "gp2"
  nat_instance_volume_size = "30"

  azs             = ["ap-northeast-1a", "ap-northeast-1c", "ap-northeast-1d"]
  private_subnets = ["10.10.0.0/24", "10.10.1.0/24", "10.10.2.0/24"]
  public_subnets  = ["10.10.10.0/24", "10.10.11.0/24", "10.10.12.0/24"]
}