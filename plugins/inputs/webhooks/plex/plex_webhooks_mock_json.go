package plex

func MediaPlayEventJSON() string {
	return `{  
		"event": "media.play",
		"user": true,
		"owner": true,
		"Account": {
			"id": 1,
			"thumb": "https://plex.tv/users/1022b120ffbaa/avatar?c=1465525047",
			"title": "elan"
		},
		"Server": {
			"title": "Office",
			"uuid": "54664a3d8acc39983675640ec9ce00b70af9cc36"
		},
		"Player": {
			"local": true,
			"publicAddress": "200.200.200.200",
			"title": "Plex Web (Safari)",
			"uuid": "r6yfkdnfggbh2bdnvkffwbms"
		},
		"Metadata": {
			"librarySectionType": "artist",
			"ratingKey": "1936545",
			"key": "/library/metadata/1936545",
			"parentRatingKey": "1936544",
			"grandparentRatingKey": "1936543",
			"guid": "com.plexapp.agents.plexmusic://gracenote/track/7572499-91016293BE6BF7F1AB2F848F736E74E5/7572500-3CBAE310D4F3E66C285E104A1458B272?lang=en",
			"librarySectionID": 1224,
			"type": "track",
			"title": "Love The One You're With",
			"grandparentKey": "/library/metadata/1936543",
			"parentKey": "/library/metadata/1936544",
			"grandparentTitle": "Stephen Stills",
			"parentTitle": "Stephen Stills",
			"summary": "",
			"index": 1,
			"parentIndex": 1,
			"ratingCount": 6794,
			"thumb": "/library/metadata/1936544/thumb/1432897518",
			"art": "/library/metadata/1936543/art/1485951497",
			"parentThumb": "/library/metadata/1936544/thumb/1432897518",
			"grandparentThumb": "/library/metadata/1936543/thumb/1485951497",
			"grandparentArt": "/library/metadata/1936543/art/1485951497",
			"addedAt": 1000396126,
			"updatedAt": 1432897518
		}
	}`
}

func MediaPlayTrailerEventJSON() string {
	return `{
		"Account": {
		  "id": 1, 
		  "thumb": "https://plex.tv/users/123456789012345/avatar?c=1234567890", 
		  "title": "username"
		}, 
		"Metadata": {
		  "addedAt": 1613901884, 
		  "art": "/library/metadata/8354/art/1614074469", 
		  "extraType": 1, 
		  "guid": "iva://api.internetvideoarchive.com/2.0/DataService/VideoAssets(409353)?lang=en&bitrates=80,212,450,750,1500,2500,5000,8000&duration=130&adaptive=1&dts=0", 
		  "index": 1, 
		  "key": "/library/metadata/8355", 
		  "originallyAvailableAt": "2019-12-19", 
		  "ratingKey": "8355", 
		  "subtype": "trailer", 
		  "summary": "", 
		  "thumb": "/library/metadata/8355/thumb/1613901884", 
		  "title": "Tenet", 
		  "type": "clip", 
		  "year": 2019
		}, 
		"Player": {
		  "local": true, 
		  "publicAddress": "11.11.11.11", 
		  "title": "Chrome", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"Server": {
		  "title": "Office", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"event": "media.play", 
		"owner": true, 
		"user": true
	  }
	`
}

func MediaStopEventJSON() string {
	return `{
		"Account": {
		  "id": 1, 
		  "thumb": "https://plex.tv/users/123456789012345/avatar?c=1234567890", 
		  "title": "username"
		}, 
		"Metadata": {
		  "Country": [
			{
			  "count": 518, 
			  "filter": "country=81", 
			  "id": 81, 
			  "tag": "USA"
			}
		  ], 
		  "Director": [
			{
			  "count": 2, 
			  "filter": "director=15153", 
			  "id": 15153, 
			  "tag": "Gore Verbinski"
			}
		  ], 
		  "Genre": [
			{
			  "count": 320, 
			  "filter": "genre=239", 
			  "id": 239, 
			  "tag": "Drama"
			}, 
			{
			  "count": 99, 
			  "filter": "genre=144", 
			  "id": 144, 
			  "tag": "Fantasy"
			}, 
			{
			  "count": 103, 
			  "filter": "genre=2960", 
			  "id": 2960, 
			  "tag": "Horror"
			}, 
			{
			  "count": 104, 
			  "filter": "genre=494", 
			  "id": 494, 
			  "tag": "Mystery"
			}, 
			{
			  "count": 102, 
			  "filter": "genre=240", 
			  "id": 240, 
			  "tag": "Sci-Fi"
			}, 
			{
			  "count": 200, 
			  "filter": "genre=241", 
			  "id": 241, 
			  "tag": "Thriller"
			}
		  ], 
		  "Producer": [
			{
			  "filter": "producer=23197", 
			  "id": 23197, 
			  "tag": "Gore Verbinski"
			}, 
			{
			  "count": 5, 
			  "filter": "producer=289", 
			  "id": 289, 
			  "tag": "Arnon Milchan"
			}, 
			{
			  "filter": "producer=23196", 
			  "id": 23196, 
			  "tag": "David Crockett"
			}
		  ], 
		  "Role": [
			{
			  "filter": "actor=23150", 
			  "id": 23150, 
			  "role": "Mr. Lockhart", 
			  "tag": "Dane DeHaan", 
			  "thumb": "http://image.tmdb.org/t/p/original/vZ7ix8UeRYgepFPFc4FIVKDLKDX.jpg"
			}, 
			{
			  "count": 7, 
			  "filter": "actor=1493", 
			  "id": 1493, 
			  "role": "Dr. Heinreich Volmer / Baron von Reichmerl", 
			  "tag": "Jason Isaacs", 
			  "thumb": "http://image.tmdb.org/t/p/original/ekOlEznbxaEfANq4ZXkpj9hVyDX.jpg"
			}, 
			{
			  "count": 2, 
			  "filter": "actor=2948", 
			  "id": 2948, 
			  "role": "Hannah von Reichmerl", 
			  "tag": "Mia Goth", 
			  "thumb": "http://image.tmdb.org/t/p/original/y3d90jZb8zsUel6viU4iIVpNBKL.jpg"
			}, 
			{
			  "count": 2, 
			  "filter": "actor=22367", 
			  "id": 22367, 
			  "role": "Deputy Director", 
			  "tag": "Adrian Schiller", 
			  "thumb": "http://image.tmdb.org/t/p/original/8SwwWk9PT5uzeHiPGtR0COeFe2W.jpg"
			}, 
			{
			  "filter": "actor=23146", 
			  "id": 23146, 
			  "role": "Victoria Watkins", 
			  "tag": "Celia Imrie", 
			  "thumb": "http://image.tmdb.org/t/p/original/tAEzs5PFVX5GsfD3KTWINnuTzMJ.jpg"
			}, 
			{
			  "filter": "actor=23160", 
			  "id": 23160, 
			  "role": "Roland Pembroke", 
			  "tag": "Harry Groener", 
			  "thumb": "http://image.tmdb.org/t/p/original/y9Z0ZIbFDwe7V7PQEMeCyTld3zM.jpg"
			}, 
			{
			  "filter": "actor=23195", 
			  "id": 23195, 
			  "role": "Frank Hill", 
			  "tag": "Tomas Norstr\u00f6m", 
			  "thumb": "http://image.tmdb.org/t/p/original/pKFRIgQreCyq2r2r1ahl26V4Cfv.jpg"
			}, 
			{
			  "filter": "actor=23164", 
			  "id": 23164, 
			  "role": "Funeral Director", 
			  "tag": "Jeff Burrell", 
			  "thumb": "http://image.tmdb.org/t/p/original/cQ6z77rmiV1iZTTORjSpFJXV5bw.jpg"
			}, 
			{
			  "filter": "actor=23162", 
			  "id": 23162, 
			  "role": "Enrico", 
			  "tag": "Ivo Nandi", 
			  "thumb": "http://image.tmdb.org/t/p/original/fa3msIf4s4kjK4dRJi6ykP44Ek2.jpg"
			}, 
			{
			  "filter": "actor=23172", 
			  "id": 23172, 
			  "role": "Pieter the Vet", 
			  "tag": "Magnus Krepper", 
			  "thumb": "http://image.tmdb.org/t/p/original/iHJ5HM0bV5UfogcOobnKJtyJklf.jpg"
			}, 
			{
			  "filter": "actor=23179", 
			  "id": 23179, 
			  "role": "Bartender", 
			  "tag": "Michael Mendl", 
			  "thumb": "http://image.tmdb.org/t/p/original/2ZxNUpT46FEPxVFJjfoLfhRifaU.jpg"
			}, 
			{
			  "filter": "actor=23149", 
			  "id": 23149, 
			  "role": "Morris", 
			  "tag": "Craig Wroe", 
			  "thumb": "http://image.tmdb.org/t/p/original/niX7llMIwMoYFjuNfgzIPYaF8EH.jpg"
			}, 
			{
			  "filter": "actor=23152", 
			  "id": 23152, 
			  "role": "Hank Green", 
			  "tag": "David Bishins", 
			  "thumb": "http://image.tmdb.org/t/p/original/9vMVIz3it1b8HRkL4BnGOFYgGpl.jpg"
			}, 
			{
			  "count": 2, 
			  "filter": "actor=7851", 
			  "id": 7851, 
			  "role": "Hollis", 
			  "tag": "Lisa Banes", 
			  "thumb": "http://image.tmdb.org/t/p/original/rXBe8lx88ESRIoKLqUMIYWpfoCq.jpg"
			}, 
			{
			  "count": 2, 
			  "filter": "actor=23144", 
			  "id": 23144, 
			  "role": "Mr. Wilson", 
			  "tag": "Carl Lumbly", 
			  "thumb": "http://image.tmdb.org/t/p/original/kJ6OPPGrWHZALsQcT8JZMQSiEbg.jpg"
			}, 
			{
			  "count": 2, 
			  "filter": "actor=9944", 
			  "id": 9944, 
			  "role": "Carl", 
			  "tag": "Jason Babinsky", 
			  "thumb": "http://image.tmdb.org/t/p/original/9Y1d7LH1NIvO6bKe6D2HZGm7Wok.jpg"
			}, 
			{
			  "filter": "actor=23165", 
			  "id": 23165, 
			  "role": "Caretaker", 
			  "tag": "Johannes Krisch", 
			  "thumb": "http://image.tmdb.org/t/p/original/28jMQgqVloCK3HAljhLHopV7aVx.jpg"
			}, 
			{
			  "count": 2, 
			  "filter": "actor=15468", 
			  "id": 15468, 
			  "role": "Mrs. Abramov", 
			  "tag": "Maggie Steed", 
			  "thumb": "http://image.tmdb.org/t/p/original/mZGUlqb27WXD6rlkpBcHXbPnf1h.jpg"
			}, 
			{
			  "filter": "actor=23140", 
			  "id": 23140, 
			  "role": "Ron Nair", 
			  "tag": "Ashok Mandanna", 
			  "thumb": "http://image.tmdb.org/t/p/original/tZtqEgjEghwzr88j4gfFlUmS5fL.jpg"
			}, 
			{
			  "filter": "actor=23156", 
			  "id": 23156, 
			  "role": "Josh", 
			  "tag": "Eric Todd", 
			  "thumb": "http://image.tmdb.org/t/p/original/j77mNkMMlM82uMVWm2ByKEO8oSH.jpg"
			}, 
			{
			  "filter": "actor=23186", 
			  "id": 23186, 
			  "role": "Lockhart's Mother", 
			  "tag": "Rebecca Street", 
			  "thumb": "http://image.tmdb.org/t/p/original/mY6xiyfA5jr2XIl2C0fM52Ap1Of.jpg"
			}, 
			{
			  "filter": "actor=23143", 
			  "id": 23143, 
			  "role": "Lockhart's Father", 
			  "tag": "Bert Tischendorf", 
			  "thumb": "http://image.tmdb.org/t/p/original/qxjrdW25tQlEnx9PKvAsDHklYvi.jpg"
			}, 
			{
			  "filter": "actor=23155", 
			  "id": 23155, 
			  "role": "9-Year-Old Lockhart", 
			  "tag": "Douglas Hamilton"
			}, 
			{
			  "filter": "actor=23170", 
			  "id": 23170, 
			  "role": "Frank", 
			  "tag": "Leonard Kunz", 
			  "thumb": "http://image.tmdb.org/t/p/original/htq6xkZ03jkwLFdgbXMbranPRvI.jpg"
			}, 
			{
			  "filter": "actor=23151", 
			  "id": 23151, 
			  "role": "Denim", 
			  "tag": "Daniel Michel", 
			  "thumb": "http://image.tmdb.org/t/p/original/lEm1ZFkTu1FZt0rOcyGdUi06MSB.jpg"
			}, 
			{
			  "filter": "actor=23194", 
			  "id": 23194, 
			  "role": "Humphrey", 
			  "tag": "Tom Flynn", 
			  "thumb": "http://image.tmdb.org/t/p/original/dqkpLiIts8ctOyzO5x0XbDYqrpJ.jpg"
			}, 
			{
			  "filter": "actor=23191", 
			  "id": 23191, 
			  "role": "Bar Girl", 
			  "tag": "Sophie Charlotte Conrad", 
			  "thumb": "http://image.tmdb.org/t/p/original/ta7dEBr9iu4fiUMTWYdu1Zr8uHF.jpg"
			}, 
			{
			  "filter": "actor=23182", 
			  "id": 23182, 
			  "role": "Bar Girl", 
			  "tag": "Natascha Lawiszus", 
			  "thumb": "http://image.tmdb.org/t/p/original/9SaoG1OPSgNhbp5A05xYinrhJLa.jpg"
			}, 
			{
			  "filter": "actor=23171", 
			  "id": 23171, 
			  "role": "Bar Girl", 
			  "tag": "Luzie Scheuritzel", 
			  "thumb": "http://image.tmdb.org/t/p/original/vlQuvw4mOpxf5CNQ9nncc5gcWUr.jpg"
			}, 
			{
			  "filter": "actor=23173", 
			  "id": 23173, 
			  "role": "Secretary", 
			  "tag": "Manon Kahle", 
			  "thumb": "http://image.tmdb.org/t/p/original/kVBJ6iHfzFVLLPulOg0qAqlrYtC.jpg"
			}, 
			{
			  "filter": "actor=23136", 
			  "id": 23136, 
			  "role": "Conductor", 
			  "tag": "Adrian Zwicker", 
			  "thumb": "http://image.tmdb.org/t/p/original/ru2aNoC3a3vvWY1ALEYXI9T2akp.jpg"
			}, 
			{
			  "filter": "actor=23184", 
			  "id": 23184, 
			  "role": "Constable", 
			  "tag": "Peter Benedict", 
			  "thumb": "http://image.tmdb.org/t/p/original/xdh1o7MRXbkdTclYjvEJHSM8RWm.jpg"
			}, 
			{
			  "filter": "actor=23167", 
			  "id": 23167, 
			  "role": "Vet's Daughter", 
			  "tag": "Julia Graefner", 
			  "thumb": "http://image.tmdb.org/t/p/original/d5FrLIZ5vCyJTwOtQ9tW4LzqlHO.jpg"
			}, 
			{
			  "filter": "actor=23183", 
			  "id": 23183, 
			  "role": "Drawing Boy", 
			  "tag": "Nino B\u00f6hlau", 
			  "thumb": "http://image.tmdb.org/t/p/original/8RN4mBvNqnYYf54CFiQT6RF54yh.jpg"
			}, 
			{
			  "filter": "actor=23139", 
			  "id": 23139, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Annette Lober", 
			  "thumb": "http://image.tmdb.org/t/p/original/orxEVtQQWfcWM5ax0AMaRKrApx8.jpg"
			}, 
			{
			  "filter": "actor=23192", 
			  "id": 23192, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Susanne Wuest", 
			  "thumb": "http://image.tmdb.org/t/p/original/zps8shfOSg9469vT5r3FRk4PHTx.jpg"
			}, 
			{
			  "filter": "actor=23158", 
			  "id": 23158, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Godehard Giese", 
			  "thumb": "http://image.tmdb.org/t/p/original/2vz3pSqdPUEtJ5xcNTWh7eZVdkj.jpg"
			}, 
			{
			  "filter": "actor=23180", 
			  "id": 23180, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Nadine Boeske", 
			  "thumb": "http://image.tmdb.org/t/p/original/sFKCSbUTPAmYGRB514mlYB2deWg.jpg"
			}, 
			{
			  "filter": "actor=23166", 
			  "id": 23166, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Judith Hoersch", 
			  "thumb": "http://image.tmdb.org/t/p/original/fHFK7L4K710I9FmRXT3nlY1aonG.jpg"
			}, 
			{
			  "filter": "actor=23138", 
			  "id": 23138, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Angelina H\u00e4ntsch", 
			  "thumb": "http://image.tmdb.org/t/p/original/h6l5XIi6piCiVuSQKx8GlGlv4SA.jpg"
			}, 
			{
			  "filter": "actor=23141", 
			  "id": 23141, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Axel Buchholz", 
			  "thumb": "http://image.tmdb.org/t/p/original/zbl2n1eOcYlHHE8bcIrSWiVky5M.jpg"
			}, 
			{
			  "filter": "actor=23153", 
			  "id": 23153, 
			  "role": "Volmer Institute Staff", 
			  "tag": "David Bredin", 
			  "thumb": "http://image.tmdb.org/t/p/original/sVgahseYW18JmE7xV5YIdmUYE4L.jpg"
			}, 
			{
			  "filter": "actor=23177", 
			  "id": 23177, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Matthias Britschgi", 
			  "thumb": "http://image.tmdb.org/t/p/original/aCgZIWDIOLzbhS27TLKQiVN3sfJ.jpg"
			}, 
			{
			  "filter": "actor=23137", 
			  "id": 23137, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Andreas Dobberkau", 
			  "thumb": "http://image.tmdb.org/t/p/original/hxw1RjSaF4QY6h1UoWsBXqRaAf2.jpg"
			}, 
			{
			  "filter": "actor=23189", 
			  "id": 23189, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Sarah Franke", 
			  "thumb": "http://image.tmdb.org/t/p/original/vlmTDjw7lj8yeXE4b5Sjic6AUfx.jpg"
			}, 
			{
			  "filter": "actor=23145", 
			  "id": 23145, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Caspar Kaeser", 
			  "thumb": "http://image.tmdb.org/t/p/original/ypSaCBscNOveJJJkbV0G7LGKGMv.jpg"
			}, 
			{
			  "filter": "actor=23190", 
			  "id": 23190, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Sebastian Kaufmane", 
			  "thumb": "http://image.tmdb.org/t/p/original/cMbaHGavHkGFYIdA9hZNdMO8g2s.jpg"
			}, 
			{
			  "filter": "actor=23175", 
			  "id": 23175, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Martin Laue", 
			  "thumb": "http://image.tmdb.org/t/p/original/eLXHS8hfwJDMzDvKpQWGMZXBqPa.jpg"
			}, 
			{
			  "filter": "actor=23142", 
			  "id": 23142, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Benedict Seifert"
			}, 
			{
			  "filter": "actor=23188", 
			  "id": 23188, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Roman Schomburg", 
			  "thumb": "http://image.tmdb.org/t/p/original/2DfABy8XDgE9L9ZbnWticaCuw8c.jpg"
			}, 
			{
			  "filter": "actor=23157", 
			  "id": 23157, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Florian Steffens", 
			  "thumb": "http://image.tmdb.org/t/p/original/n7lzrYXuq9rRGW5IiHEctLisldE.jpg"
			}, 
			{
			  "filter": "actor=23193", 
			  "id": 23193, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Sven Timmreck", 
			  "thumb": "http://image.tmdb.org/t/p/original/fYj8BrPO9L4T1gp9BObXjd05QvK.jpg"
			}, 
			{
			  "filter": "actor=23159", 
			  "id": 23159, 
			  "role": "Volmer Institute Staff", 
			  "tag": "Hanna Wollschlaeger"
			}, 
			{
			  "filter": "actor=23163", 
			  "id": 23163, 
			  "role": "Volmer Institute Patient", 
			  "tag": "Jef Bayonne", 
			  "thumb": "http://image.tmdb.org/t/p/original/eWGAzJsksajKRM6olTVci6l8iGX.jpg"
			}, 
			{
			  "filter": "actor=23169", 
			  "id": 23169, 
			  "role": "Volmer Institute Patient", 
			  "tag": "Klaus Burckhardt"
			}, 
			{
			  "filter": "actor=23161", 
			  "id": 23161, 
			  "role": "Volmer Institute Patient", 
			  "tag": "Helmut Boelling"
			}, 
			{
			  "filter": "actor=23187", 
			  "id": 23187, 
			  "role": "Volmer Institute Patient", 
			  "tag": "Renate Prack"
			}, 
			{
			  "filter": "actor=23178", 
			  "id": 23178, 
			  "role": "Volmer Institute Patient", 
			  "tag": "Matthias Rick"
			}, 
			{
			  "filter": "actor=23185", 
			  "id": 23185, 
			  "role": "Volmer Institute Patient", 
			  "tag": "Peter Simlinger", 
			  "thumb": "http://image.tmdb.org/t/p/original/1Vs8guayqsiJgWiKqHifS8ZmI1K.jpg"
			}, 
			{
			  "filter": "actor=23168", 
			  "id": 23168, 
			  "role": "Volmer Institute Patient", 
			  "tag": "Jutta Marina Von Brunkau"
			}, 
			{
			  "count": 4, 
			  "filter": "actor=17109", 
			  "id": 17109, 
			  "role": "Doctor Barlock (uncredited)", 
			  "tag": "Johnny Otto", 
			  "thumb": "http://image.tmdb.org/t/p/original/i5LOfr6GLH6hvBdCoMyKY457IQZ.jpg"
			}, 
			{
			  "filter": "actor=23181", 
			  "id": 23181, 
			  "role": "Nurse (uncredited)", 
			  "tag": "Natalia Bobrich", 
			  "thumb": "http://image.tmdb.org/t/p/original/shi8hbeikCZ7Hgwess3E9yKC76D.jpg"
			}, 
			{
			  "count": 2, 
			  "filter": "actor=17247", 
			  "id": 17247, 
			  "role": "Wall Street Banker (uncredited)", 
			  "tag": "Alexander Yassin", 
			  "thumb": "http://image.tmdb.org/t/p/original/gnB8XCSUbGqi1JCqj7IR3tYedUK.jpg"
			}, 
			{
			  "filter": "actor=23148", 
			  "id": 23148, 
			  "role": "Wallstreet Broker (uncredited)", 
			  "tag": "Chris Theisinger", 
			  "thumb": "http://image.tmdb.org/t/p/original/l5eM8YTOFbVBQFgAEfYn8oiDZ08.jpg"
			}, 
			{
			  "filter": "actor=23154", 
			  "id": 23154, 
			  "role": "Orderly (uncredited)", 
			  "tag": "David S\u00e1nchez Calvo"
			}, 
			{
			  "count": 2, 
			  "filter": "actor=17285", 
			  "id": 17285, 
			  "role": "Ron NYC Driver (uncredited)", 
			  "tag": "Earl Vincent Sherwood II"
			}, 
			{
			  "filter": "actor=23147", 
			  "id": 23147, 
			  "role": "Wedding Guest (uncredited)", 
			  "tag": "Chris Huszar"
			}, 
			{
			  "filter": "actor=23174", 
			  "id": 23174, 
			  "role": "Wedding Guest (uncredited)", 
			  "tag": "Marko Buzin"
			}, 
			{
			  "filter": "actor=23176", 
			  "id": 23176, 
			  "role": "(voice) (uncredited)", 
			  "tag": "Matt Lindquist", 
			  "thumb": "http://image.tmdb.org/t/p/original/az4dpSBsDbQvoDxRLIYJal6bO5z.jpg"
			}
		  ], 
		  "Similar": [
			{
			  "count": 9, 
			  "filter": "similar=11451", 
			  "id": 11451, 
			  "tag": "The Autopsy of Jane Doe"
			}, 
			{
			  "count": 9, 
			  "filter": "similar=10594", 
			  "id": 10594, 
			  "tag": "Life"
			}, 
			{
			  "count": 12, 
			  "filter": "similar=7503", 
			  "id": 7503, 
			  "tag": "It Comes at Night"
			}, 
			{
			  "count": 5, 
			  "filter": "similar=9772", 
			  "id": 9772, 
			  "tag": "Raw"
			}, 
			{
			  "count": 14, 
			  "filter": "similar=9491", 
			  "id": 9491, 
			  "tag": "Lights Out"
			}, 
			{
			  "count": 2, 
			  "filter": "similar=9227", 
			  "id": 9227, 
			  "tag": "The Belko Experiment"
			}, 
			{
			  "filter": "similar=23199", 
			  "id": 23199, 
			  "tag": "Morgan"
			}, 
			{
			  "count": 2, 
			  "filter": "similar=9774", 
			  "id": 9774, 
			  "tag": "The Neon Demon"
			}, 
			{
			  "count": 4, 
			  "filter": "similar=11445", 
			  "id": 11445, 
			  "tag": "Blair Witch"
			}, 
			{
			  "count": 7, 
			  "filter": "similar=11455", 
			  "id": 11455, 
			  "tag": "The Visit"
			}, 
			{
			  "filter": "similar=23198", 
			  "id": 23198, 
			  "tag": "Before I Wake"
			}, 
			{
			  "count": 9, 
			  "filter": "similar=7505", 
			  "id": 7505, 
			  "tag": "The Witch"
			}, 
			{
			  "count": 11, 
			  "filter": "similar=11446", 
			  "id": 11446, 
			  "tag": "Don't Breathe"
			}, 
			{
			  "count": 15, 
			  "filter": "similar=7000", 
			  "id": 7000, 
			  "tag": "Mother!"
			}, 
			{
			  "count": 2, 
			  "filter": "similar=21355", 
			  "id": 21355, 
			  "tag": "The Void"
			}, 
			{
			  "count": 4, 
			  "filter": "similar=9228", 
			  "id": 9228, 
			  "tag": "The Girl with All the Gifts"
			}, 
			{
			  "count": 7, 
			  "filter": "similar=9493", 
			  "id": 9493, 
			  "tag": "Ouija: Origin of Evil"
			}, 
			{
			  "count": 7, 
			  "filter": "similar=11448", 
			  "id": 11448, 
			  "tag": "It Follows"
			}, 
			{
			  "count": 5, 
			  "filter": "similar=9495", 
			  "id": 9495, 
			  "tag": "The Bye Bye Man"
			}, 
			{
			  "count": 7, 
			  "filter": "similar=10592", 
			  "id": 10592, 
			  "tag": "Alien: Covenant"
			}
		  ], 
		  "Writer": [
			{
			  "filter": "writer=23135", 
			  "id": 23135, 
			  "tag": "Justin Haythe"
			}
		  ], 
		  "addedAt": 1555298191, 
		  "art": "/library/metadata/564/art/1555495001", 
		  "audienceRating": 4.1, 
		  "audienceRatingImage": "rottentomatoes://image.rating.spilled", 
		  "contentRating": "R", 
		  "duration": 8760000, 
		  "guid": "com.plexapp.agents.imdb://tt4731136?lang=en", 
		  "key": "/library/metadata/564", 
		  "librarySectionID": 2, 
		  "librarySectionKey": "/library/sections/2", 
		  "librarySectionTitle": "Movies", 
		  "librarySectionType": "movie", 
		  "originallyAvailableAt": "2017-02-15", 
		  "primaryExtraKey": "/library/metadata/3598", 
		  "rating": 4.3, 
		  "ratingImage": "rottentomatoes://image.rating.rotten", 
		  "ratingKey": "564", 
		  "studio": "20th Century Fox", 
		  "summary": "An ambitious young executive is sent to retrieve his company's CEO from an idyllic but mysterious \"wellness center\" at a remote location in the Swiss Alps but soon suspects that the spa's miraculous treatments are not what they seem.", 
		  "tagline": "There is a Cure", 
		  "thumb": "/library/metadata/564/thumb/1555495001", 
		  "title": "EVO", 
		  "type": "movie", 
		  "updatedAt": 1555495001, 
		  "year": 2017
		}, 
		"Player": {
		  "local": true, 
		  "publicAddress": "11.11.11.11", 
		  "title": "Chrome", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"Server": {
		  "title": "Office", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"event": "media.stop", 
		"owner": true, 
		"user": true
	  }
	  `
}

func MediaPhotoEventJSON() string {
	return `{
		"Account": {
			"id": 1, 
			"thumb": "https://plex.tv/users/123456789012345/avatar?c=1234567890", 
			"title": "username"
		}, 
		"Metadata": {
			"addedAt": 1614209252, 
			"createdAtAccuracy": "local", 
			"createdAtTZOffset": "-28800", 
			"guid": "local://8369", 
			"index": 1, 
			"key": "/library/metadata/8369", 
			"librarySectionID": 4, 
			"librarySectionKey": "/library/sections/4", 
			"librarySectionTitle": "Photos", 
			"librarySectionType": "photo", 
			"originallyAvailableAt": "2021-02-24", 
			"ratingKey": "8369", 
			"summary": "", 
			"thumb": "/library/metadata/8369/thumb/1614209253", 
			"title": "Screen Shot 2021-02-24 at 3.26.14 PM", 
			"type": "photo", 
			"updatedAt": 1614209253, 
			"year": 2021
		}, 
		"Player": {
			"local": true, 
			"publicAddress": "11.11.11.11", 
			"title": "Chrome", 
			"uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"Server": {
			"title": "Office", 
			"uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"event": "media.play", 
		"owner": true, 
		"user": true
	}
	`
}

func MediaStreamingMovieEvent() string {
	return `{
		"Account": {
		  "id": 2534941, 
		  "title": "username"
		}, 
		"Metadata": {
		  "addedAt": 1612166400, 
		  "art": "https://image.tmdb.org/t/p/original/eJ6kPtua2rEvjxHdCtnIobXJZSf.jpg", 
		  "attributionLogo": "https://provider-static.plex.tv/vod/partners/logos/crackle.png", 
		  "audienceRating": 9.3, 
		  "audienceRatingImage": "rottentomatoes://image.rating.upright", 
		  "banner": "https://assets.fanart.tv/fanart/movies/829/moviebanner/chinatown-5318abafa0ef1.jpg", 
		  "contentRating": "R", 
		  "duration": 7829994, 
		  "guid": "plex://movie/5d7768273c3c2a001fbcb51b", 
		  "indirect": true, 
		  "key": "/library/metadata/6015de7207408f002dcee048", 
		  "originallyAvailableAt": "1974-06-20T00:00:00Z", 
		  "primaryExtraKey": "/library/metadata/6015de7207408f002dcee048/extras/iva-65", 
		  "rating": 9.9, 
		  "ratingCount": 1933, 
		  "ratingImage": "rottentomatoes://image.rating.ripe", 
		  "ratingKey": "6015de7207408f002dcee048", 
		  "studio": "Paramount", 
		  "thumb": "https://image.tmdb.org/t/p/original/7ljxmU9L0ugUvyyg0GBkFaarjEq.jpg", 
		  "title": "Chinatown", 
		  "type": "movie", 
		  "viewCount": 0, 
		  "viewOffset": 0, 
		  "year": 1974
		}, 
		"Player": {
		  "local": false, 
		  "publicAddress": "11.11.11.11", 
		  "title": "Plex Web (Chrome)", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"Server": {
		  "title": "Vod", 
		  "uuid": "tv.plex.provider.vod"
		}, 
		"event": "media.stop", 
		"owner": false, 
		"user": true
	  }
	  `
}

func MediaScrobbleEvent() string {
	return `{
		"Account": {
		  "id": 1, 
		  "thumb": "https://plex.tv/users/123456789012345/avatar?c=1234567890", 
		  "title": "username"
		}, 
		"Metadata": {
		  "Channel": [
			{
			  "count": 2, 
			  "filter": "channel=33160", 
			  "id": 33160, 
			  "tag": "2.3 KTVUDT3 (Movies!)"
			}, 
			{
			  "filter": "channel=33315", 
			  "id": 33315, 
			  "tag": "66.4 KFSFDT4 (Grit)"
			}
		  ], 
		  "Field": [
			{
			  "locked": true, 
			  "name": "thumb"
			}, 
			{
			  "locked": true, 
			  "name": "art"
			}
		  ], 
		  "Genre": [
			{
			  "count": 19, 
			  "filter": "genre=287", 
			  "id": 287, 
			  "tag": "Western"
			}
		  ], 
		  "addedAt": 1614242259, 
		  "art": "/library/metadata/8372/art/1614242261", 
		  "contentRating": "TVPG", 
		  "duration": 7200000, 
		  "guid": "plex://movie/5fc690d7578684002e68e65d", 
		  "key": "/library/metadata/8372", 
		  "lastViewedAt": 1614273229, 
		  "librarySectionID": 2, 
		  "librarySectionKey": "/library/sections/2", 
		  "librarySectionTitle": "Movies", 
		  "librarySectionType": "movie", 
		  "oneShot": "1", 
		  "originallyAvailableAt": "1957-01-01", 
		  "ratingKey": "8372", 
		  "summary": "Dan Evans (Van Heflin), a drought-plagued Arizona rancher, volunteers to take captured stagecoach robber and murderer Ben Wade (Glenn Ford) from Bisbee to Contention City, where the criminal will be put aboard the 3:10 train to Yuma for his trial. Accompanied only by the town drunk, Alex Potter (Henry Jones), Dan battles Wade's henchman (Richard Jaeckel), the murder victim's revenge-minded brother, and the temptation of the large bribe Wade offers in exchange for his freedom.", 
		  "thumb": "/library/metadata/8372/thumb/1614242261", 
		  "title": "3:10 to Yuma", 
		  "type": "movie", 
		  "updatedAt": 1614242261, 
		  "viewCount": 1, 
		  "year": 1957
		}, 
		"Player": {
		  "local": true, 
		  "publicAddress": "11.11.11.11", 
		  "title": "Chrome", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"Server": {
		  "title": "Office", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"event": "media.scrobble", 
		"owner": true, 
		"user": true
	  }
	  `
}

func MediaMusicEventJSON() string {
	return `{
		"Account": {
		  "id": 1, 
		  "thumb": "https://plex.tv/users/123456789012345/avatar?c=1234567890", 
		  "title": "username"
		}, 
		"Metadata": {
		  "Mood": [
			{
			  "filter": "mood=33368", 
			  "id": 33368, 
			  "tag": "Lively"
			}, 
			{
			  "filter": "mood=33369", 
			  "id": 33369, 
			  "tag": "Marching"
			}, 
			{
			  "filter": "mood=33370", 
			  "id": 33370, 
			  "tag": "Pulsing"
			}, 
			{
			  "filter": "mood=33371", 
			  "id": 33371, 
			  "tag": "Strong"
			}, 
			{
			  "filter": "mood=33372", 
			  "id": 33372, 
			  "tag": "Swinging"
			}
		  ], 
		  "addedAt": 1614031681, 
		  "art": "/library/metadata/8363/art/1614031704", 
		  "grandparentArt": "/library/metadata/8363/art/1614031704", 
		  "grandparentGuid": "plex://artist/5d07bbfc403c6402904a5ec9", 
		  "grandparentKey": "/library/metadata/8363", 
		  "grandparentRatingKey": "8363", 
		  "grandparentThumb": "/library/metadata/8363/thumb/1614031704", 
		  "grandparentTitle": "Various Artists", 
		  "guid": "plex://track/5d07cdb2403c640290f4b12a", 
		  "index": 6, 
		  "key": "/library/metadata/8365", 
		  "librarySectionID": 3, 
		  "librarySectionKey": "/library/sections/3", 
		  "librarySectionTitle": "Music", 
		  "librarySectionType": "artist", 
		  "originalTitle": "Britney Spears", 
		  "parentGuid": "plex://album/5d07c886403c640290c06f67", 
		  "parentIndex": 2, 
		  "parentKey": "/library/metadata/8364", 
		  "parentRatingKey": "8364", 
		  "parentThumb": "/library/metadata/8364/thumb/1614031707", 
		  "parentTitle": "RTL: I Like the 90's", 
		  "parentYear": 2014, 
		  "ratingCount": 460901, 
		  "ratingKey": "8365", 
		  "summary": "", 
		  "thumb": "/library/metadata/8364/thumb/1614031707", 
		  "title": "Oops!\u2026I Did It Again (album version)", 
		  "type": "track", 
		  "updatedAt": 1614031707
		}, 
		"Player": {
		  "local": true, 
		  "publicAddress": "11.11.11.11", 
		  "title": "Chrome", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"Server": {
		  "title": "Office", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"event": "media.play", 
		"owner": true, 
		"user": true
	  }
	  `
}

func RateEventJSON() string {
	return `{
		"Account": {
		  "id": 1, 
		  "thumb": "https://plex.tv/users/123456789012345/avatar?c=1234567890", 
		  "title": "username"
		}, 
		"Metadata": {
		  "Channel": [
			{
			  "count": 3, 
			  "filter": "channel=33160", 
			  "id": 33160, 
			  "tag": "2.3 KTVUDT3 (Movies!)"
			}
		  ], 
		  "Field": [
			{
			  "locked": true, 
			  "name": "thumb"
			}, 
			{
			  "locked": true, 
			  "name": "art"
			}
		  ], 
		  "Genre": [
			{
			  "count": 322, 
			  "filter": "genre=239", 
			  "id": 239, 
			  "tag": "Drama"
			}
		  ], 
		  "addedAt": 1614275328, 
		  "art": "/library/metadata/8373/art/1614275330", 
		  "contentRating": "TVPG", 
		  "duration": 8400000, 
		  "guid": "plex://movie/5fc6912f1a65df002dd28ab7", 
		  "key": "/library/metadata/8373", 
		  "lastRatedAt": 1614284551, 
		  "librarySectionID": 2, 
		  "librarySectionKey": "/library/sections/2", 
		  "librarySectionTitle": "Movies", 
		  "librarySectionType": "movie", 
		  "oneShot": "1", 
		  "originallyAvailableAt": "1956-01-01", 
		  "ratingKey": "8373", 
		  "summary": "Broke and without work, newspaper reporter Eddie Willis (Humphrey Bogart) agrees to work for the corrupt boxing promoter Nick Benko (Rod Steiger) to help hype his new boxer, Toro Moreno (Mike Lane). While Toro is beastly in appearance, he has no actual boxing talent, and all his fights are fixed. When Toro gets a shot at the title against the brutal Buddy Brannen (Max Baer), Willis is faced with the tough decision of whether or not to tell Toro that his entire career is a sham.", 
		  "thumb": "/library/metadata/8373/thumb/1614275330", 
		  "title": "The Harder They Fall", 
		  "titleSort": "Harder They Fall", 
		  "type": "movie", 
		  "updatedAt": 1614275330, 
		  "userRating": 10.0, 
		  "year": 1956
		}, 
		"Player": {
		  "local": true, 
		  "publicAddress": "11.11.11.11", 
		  "title": "Chrome", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"Server": {
		  "title": "Office", 
		  "uuid": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		}, 
		"event": "media.rate", 
		"owner": true, 
		"rating": "10", 
		"user": true
	  }
	  `
}
