variable "s3_bucket" {
  description = "s3 bucket"
  type        = string
}
variable "region" {
  description = "Region to create resources"
  type        = string
}

# database config
variable "db_name" {
  description = "database name"
  type        = string
}
variable "db_username" {
  description = "database username"
  type        = string
}
variable "db_password" {
  description = "database password"
  type        = string
}

# ec2 vm
variable "instance_type" {
  description = "ec2 instance type"
  type        = string
}
variable "instance_ami" {
  description = "ec2 instance ami"
  type        = string
}
variable "key_name" {
  description = "ec2 key pair"
  type        = string
}


