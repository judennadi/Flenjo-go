resource "aws_elasticache_subnet_group" "cache-subnet-group" {
  name       = "dev-cache-subnets"
  subnet_ids = [aws_subnet.dev-private-data-subnet.id, ]

  tags = {
    Name = "dev-cache-subnets"
  }
}

resource "aws_elasticache_cluster" "cache_instance" {
  cluster_id           = "dev-cache"
  engine               = "redis"
  engine_version       = "7.0"
  parameter_group_name = "default.redis7"
  node_type            = "cache.t2.micro"
  port                 = 6379
  num_cache_nodes      = 1
  security_group_ids   = [aws_security_group.cache-security-group.id]
  subnet_group_name    = aws_elasticache_subnet_group.cache-subnet-group.name
}
