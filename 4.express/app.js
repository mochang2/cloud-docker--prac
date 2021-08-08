var express = require("express"); // express library
var app = express();

app.get("/", function(req, res){
    res.send("Hello Express");
});

app.listen(9000, () => {
    console.log("My REST API running on port 9000!");
})