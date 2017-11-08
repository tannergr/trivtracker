var map;
var iconBase = 'https://maps.google.com/mapfiles/kml/shapes/';
var mapstart = {};

var markers = [];
mapstart.lat=49.267;
mapstart.long=-123.1455;

function initAutocomplete() {
  map = new google.maps.Map(document.getElementById('map'), {
  center: {lat: mapstart.lat, lng: mapstart.long},
    zoom: 13,
    mapTypeId: 'roadmap',
    //styles: mapstyle,
    zoomControlOptions: {
                    style: google.maps.ZoomControlStyle.SMALL,
                    position: google.maps.ControlPosition.RIGHT_BOTTOM
                },
                linksControl: false,
                panControl: false,
                addressControl: false,
                enableCloseButton: false,
                fullscreenControl: false,
                enableCloseButton: false,
                streetViewControl: false
  });

  // Create the search box and link it to the UI element.
  var input = document.getElementById('location');
  var searchBox = new google.maps.places.SearchBox(input);

  // Bias the SearchBox results towards current map's viewport.
  map.addListener('bounds_changed', function() {
    searchBox.setBounds(map.getBounds());
  });

  // Listen for the event fired when the user selects a prediction and retrieve
  // more details for that place.
  searchBox.addListener('places_changed', function() {
    var places = searchBox.getPlaces();

    if (places.length == 0) {
      return;
    }

    // Clear out the old markers.
    markers.forEach(function(marker) {
      marker.setMap(null);
    });

    markers = [];
    var iconBase = 'https://maps.google.com/mapfiles/kml/shapes/';
    UpdateMarkers();


    // For each place, get the icon, name and location.
    var bounds = new google.maps.LatLngBounds();
    places.forEach(function(place) {
      if (!place.geometry) {
        console.log("Returned place contains no geometry");
        return;
      }
      var icon = {
        url: place.icon,
        size: new google.maps.Size(71, 71),
        origin: new google.maps.Point(0, 0),
        anchor: new google.maps.Point(17, 34),
        scaledSize: new google.maps.Size(25, 25)
      };

      // Create a marker for each place.
      markers.push(new google.maps.Marker({
        map: map,
        icon: icon,
        title: place.name,
        position: place.geometry.location
      }));

      if (place.geometry.viewport) {
        // Only geocodes have viewport.
        bounds.union(place.geometry.viewport);
      } else {
        bounds.extend(place.geometry.location);
      }
    });
    map.fitBounds(bounds);
  });
  window.onload = ()=>{

    var urlstring = "./places?lat=" + mapstart.lat + "&long=" + mapstart.long + "$day=" + $(".Day").val();
    $.getJSON(urlstring, null, (r)=>{
      r.forEach((place)=>{
        var marker = new google.maps.Marker({
          icon: './trivia.png',
          position: new google.maps.LatLng(place.lat, place.long),
          map: map
        });
        var infowindow = new google.maps.InfoWindow({
          content: infoWindowBuilder(place)
        });
        marker.addListener('click', function() {
          infowindow.open(map, marker);
        });
        markers.push(marker);
      })});

      $("#Day").change(()=>{
        console.log("day change");
        UpdateMarkers();
      });

      $("#event").change(()=>{
        console.log("event change");
        UpdateMarkers();
      });
  }

}

function UpdateMarkers(){

  // Clear out the old markers.
  markers.forEach(function(marker) {
    marker.setMap(null);
  });

  markers = [];
  var urlstring = "./places?lat=" + mapstart.lat + "&long=" + mapstart.long + "&day=" + $("#Day").val() + "&type=" + $("#event").val();
  console.log(urlstring);
  $.getJSON(urlstring, null, (r)=>{
    r.forEach((place)=>{
      var marker = new google.maps.Marker({
        icon: './trivia.png',
        position: new google.maps.LatLng(place.lat, place.long),
        map: map
      });
      var infowindow = new google.maps.InfoWindow({
        content: infoWindowBuilder(place)
      });
      marker.addListener('click', function() {
        infowindow.open(map, marker);
      });
      markers.push(marker);
    })});
}

function infoWindowBuilder(place){
  var str = "";
  str += "<h1>" + place.barname + "</h1>";
  str += "<h2>" + place.etype + "</h2>";
  str += "<p>" + place.comments + "</p>";
  return str;
}
