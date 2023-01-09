const conn = new Mongo()
const db = conn.getDB("gohcms")

db.createCollection('articles', {
	validator: {
		$jsonSchema: {
			required: ['titleID', 'title', 'content', 'date', 'tags', 'online'],
			properties: {
				_id: {
					bsonType: 'objectId'
				},
				titleID: {
					bsonType: 'string',
					pattern: '^[-a-z]+$'
				},
				title: {
					bsonType: 'string'
				},
				content: {
					bsonType: 'object'
				},
				date: {
					bsonType: 'number'
				},
				tags: {
					bsonType: 'array'
				},
				online: {
					bsonType: 'bool'
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