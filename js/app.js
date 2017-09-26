var forward_timeout, move_forward = $('#move-forward');
var left_timeout, move_left = $('#move-left');
var right_timeout, move_right = $('#move-right');
var backward_timeout, move_backward = $('#move-backward');
var interval_time = 100;

//Forward movement interval
move_forward.mousedown(function(){
    forward_timeout = setInterval(function(){
        submit_motion("forward");
    }, interval_time);

    return false;
});

//Left movement interval
move_left.mousedown(function(){
    left_timeout = setInterval(function(){
        submit_motion("left");
    }, interval_time);

    return false;
});

//Right movement interval
move_right.mousedown(function(){
    right_timeout = setInterval(function(){
        submit_motion("right");
    }, interval_time);

    return false;
});

//Backward movement interval
move_backward.mousedown(function(){
    backward_timeout = setInterval(function(){
        submit_motion("backward");
    }, interval_time);

    return false;
});

$(document).mouseup(function(){
    clearInterval(forward_timeout);
    clearInterval(left_timeout);
    clearInterval(right_timeout);
    clearInterval(backward_timeout);
    return false;
});

function submit_motion(direction){
  var motion = {direction: direction};
  $.ajax({
   type:"POST",
   url: "/api/motion",
   data: JSON.stringify(motion),
   processData: false,
   });
}

function toggle_lights(){
  $.ajax({
   type:"GET",
   url: "/api/lights",
   processData: false,
   });
}
