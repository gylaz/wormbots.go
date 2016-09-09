(function() {
  var source = new EventSource("/world");
  var canvas = document.getElementById("world");
  var context = canvas.getContext("2d");
  var opacity = "1";
  var rgbRange = "240, 30, 8";

  source.onopen = function() { console.log("connection opened"); };
  source.onclose = function() { console.log("connection closed"); };
  source.onmessage = function(e) {
    var worms = e.data;
    var worms = JSON.parse(e.data);

    context.clearRect(0, 0, canvas.width, canvas.height);

    worms.forEach(function(worm) {
      // console.log(worm);
      context.fillStyle = "rgba(" + rgbRange + ", " + opacity + ")";
      context.fillRect(worm.x, worm.y, 1, 1);
    });
  };
})();
