(function() {
  var source = new EventSource("/world");
  var canvas = document.getElementById("world");
  var context = canvas.getContext("2d");

  source.onopen = function() { console.log("connection opened"); };
  source.onclose = function() { console.log("connection closed"); };
  source.onmessage = function(e) {
    var worms = e.data;

    context.clearRect(0, 0, canvas.width, canvas.height);
    context.font = "32px serif";
    context.fillText(worms, 100, 40);
  };
})();
