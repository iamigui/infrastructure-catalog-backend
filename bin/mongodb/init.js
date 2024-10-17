db = db.getSiblingDB("catalog"); // Use your desired database name

db.projects.insertMany([
  {
    id: "1",
    name: "Project A",
    description: "Description for Project A",
    jsonData: JSON.stringify({ key1: "value1", key2: "value2" }), // Store as a string
  },
  {
    id: "2",
    name: "Project B",
    description: "Description for Project B",
    jsonData: JSON.stringify({ key1: "value2", key2: "value3" }), // Store as a string
  },
  {
    id: "3",
    name: "Project C",
    description: "Description for Project C",
    jsonData: JSON.stringify({ key1: "value3" }), // Store as a string
  },
]);

db.createUser ({
  user: "writer",
  pwd: "writer",
  roles: [{ role: "readWrite", db: "catalog" }]
});