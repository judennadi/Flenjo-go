resource "aws_instance" "client_ec2" {
  ami                         = var.instance_ami
  instance_type               = var.instance_type
  key_name                    = var.key_name
  subnet_id                   = aws_subnet.dev-public-client-subnet.id
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.client-security-group.id, aws_security_group.ssh-security-group.id]

  tags = {
    Name = "dev-client"
  }
}

resource "aws_instance" "server_ec2" {
  ami                         = var.instance_ami
  instance_type               = var.instance_type
  key_name                    = var.key_name
  subnet_id                   = aws_subnet.dev-private-server-subnet.id
  associate_public_ip_address = false
  vpc_security_group_ids      = [aws_security_group.server-security-group.id, aws_security_group.ssh-security-group.id]

  tags = {
    Name = "dev-server"
  }
}
