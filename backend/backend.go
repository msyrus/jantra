package backend

import "github.com/streadway/amqp"

// AMQP is the AMQP backend of the server
type AMQP struct{ *amqp.Connection }
