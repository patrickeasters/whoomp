var inTime = 47.25
var outTime = 54.5

var myWhoomps

$( document ).ready(function() {
    $.get( '/api/count', function( data ) {
      $('#total').html( Number(data) - myWhoomps );
    });
    addWhoomp();
});

// Load YouTube API async
var tag = document.createElement('script');
tag.src = "https://www.youtube.com/iframe_api";
var firstScriptTag = document.getElementsByTagName('script')[0];
firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);

// Youtube player API
var player;
function onYouTubeIframeAPIReady() {
        player = new YT.Player('player', {
          height: '405',
          width: '720',
          videoId: 'Z-FPimCmbX8',
          events: {
            'onReady': onPlayerReady,
            'onStateChange': onPlayerStateChange
          }
        });
        }

function onPlayerReady(event) {
  goToStart(event);
}

function onPlayerStateChange(event) {
  var tm = player.getCurrentTime();
  if(tm < inTime) {
    goToStart(event);
  }
  if (event.data == YT.PlayerState.PLAYING) {
    var iv = setInterval(checkEnd,100);
  }
  if (event.data == YT.PlayerState.PAUSED) {
    clearInterval(iv);
  }

  function checkEnd() {
    var ct = player.getCurrentTime();
    if (ct >= outTime) {
      goToStart(event);
      addWhoomp();
    }
  }

  if (event.data == YT.PlayerState.ENDED) {
    goToStart(event);
    addWhoomp();
  }
}

function goToStart(event) {
  event.target.seekTo(inTime,1);
  player.playVideo();
}

function addWhoomp() {
  myWhoomps = Number( localStorage.getItem('whoomps') );
  localStorage.setItem('whoomps', ++myWhoomps)
  $('#whoomps').html( myWhoomps );
  if (myWhoomps > 1){
    $('#time_plural').html('s');
  }
  $.post('/api/incr');
}
