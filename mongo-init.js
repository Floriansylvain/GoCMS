const conn = new Mongo()
const db = conn.getDB("gohcms")

db.createCollection('articles', {
	validator: {
		$jsonSchema: {
			required: ['content', 'date', 'id_name'],
			properties: {
				_id: {
					bsonType: 'objectId'
				},
				content: {
					bsonType: 'object'
				},
				date: {
					bsonType: 'number'
				},
				id_name: {
					bsonType: 'string',
					pattern: '^.+$'
				}
			}
		}
	}
})

db.createCollection('users', {
	validator: {
		$jsonSchema: {
			required: ['email', 'password'],
			properties: {
				_id: {
					bsonType: 'objectId'
				},
				email: {
					bsonType: 'string'
				},
				password: {
					bsonType: 'string'
				}
			}
		}
	}
})

db.users.insertOne({
	email: "root",
	password: "root"
})