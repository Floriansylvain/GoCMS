conn = new Mongo()
db = conn.getDB("gohcms")

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
