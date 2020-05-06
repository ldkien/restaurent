package entity

type GlobalConfig struct {
	Cassandra   Cassandra `json:"cassandra"`
	LoginClient []string  `json:"loginClient"`
	LoginPort   string    `json:"loginPort"`
}

type Cassandra struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Keyspace string `json:"keyspace"`
}
