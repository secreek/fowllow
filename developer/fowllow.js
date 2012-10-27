var http = new XMLHttpRequest();

function put_url(name, meeting, url) {
  http.open('PUT', 'http://localhost:4567/' + name + "/" + meeting + "?url=" + url, false);
  http.send('url=' + url);
}

function get_url(name, meeting, callback) {
  http.open('GET', 'http://localhost:4567/' + name + "/" + meeting, true);
  http.onreadystatechange=function() {
    if (http.readyState==4 && http.status==200) {
      if (callback) {
        callback(http.responseText);
      };
    }
  }
  http.send();
}
