package entity

type GlobalConfig struct {
	Cassandra   Cassandra `json:"cassandra"`
	LoginClient []string  `json:"loginClient"`
	OrderClient []string  `json:"orderClient"`
	LoginPort   string    `json:"loginPort"`
	OrderPort   string    `json:"orderPort"`
}

type Cassandra struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Keyspace string `json:"keyspace"`
}
