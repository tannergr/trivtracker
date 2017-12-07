var map;
var iconBase = 'https://maps.google.com/mapfiles/kml/shapes/';
var mapstart = {};

var markers = [];
mapstart.lat=49.267;
mapstart.long=-123.1455;
var lastwindow;
var service;

var userposition = {};
userposition.set = false;
function initAutocomplete() {
  window.onload =()=>{
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
    service = new google.maps.places.PlacesService(map);
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
      map.setCenter(searchBox.getPlaces()[0].geometry.location);
      UpdateMarkers();
    });
    UpdateMarkers();
    init();
  }
}
function init(){
    var welcometextbox = document.getElementById('startlocation');
    var welcomesearchBox = new google.maps.places.SearchBox(welcometextbox);


    $("#Day").change(()=>{
      console.log("day change");
      UpdateMarkers();
    });

    $("#event").change(()=>{
      console.log("event change");
      UpdateMarkers();
    });
    $('.ui.small.modal').modal('show');
    $('#uselocationbutton').click(()=>{
      navigator.geolocation.getCurrentPosition(setLocation);
    });
    $('#searchbutton').click(()=>{
      //set center
      if(userposition.set){
        map.setCenter({lat:userposition.lat, lng:userposition.long});
      }
      else{
        if(welcomesearchBox.getPlaces())
          map.setCenter(welcomesearchBox.getPlaces()[0].geometry.location);
      }
      UpdateMarkers(true);
    });
}


function setLocation(pos){
   var crd = pos.coords;
   console.log('Your current position is:');
   console.log(`Latitude : ${crd.latitude}`);
   console.log(`Longitude: ${crd.longitude}`);
   console.log(`More or less ${crd.accuracy} meters.`);
   $('#uselocationbutton').toggleClass("active");
   userposition.set = true;
   userposition.lat = crd.latitude;
   userposition.long = crd.longitude;
}
function UpdateMarkers(isWelcome){

  // Clear out the old markers.
  markers.forEach(function(marker) {
    marker.setMap(null);
  });

  markers = [];
  openwindows = [];
  if(isWelcome)
    var urlstring = "./places?lat=" + mapstart.lat + "&long=" + mapstart.long + "&day=" + $("#startDay").val() + "&type=" + $("#startevent").val();
  else
    var urlstring = "./places?lat=" + mapstart.lat + "&long=" + mapstart.long + "&day=" + $("#Day").val() + "&type=" + $("#event").val();
  console.log(urlstring);
  $.getJSON(urlstring, null, (r)=>{
    r.forEach((place)=>{
      var marker = new google.maps.Marker({
        icon: getIcon(place),
        position: new google.maps.LatLng(place.lat, place.long),
        map: map
      });
      var infowindow = new google.maps.InfoWindow({
        content: infoWindowBuilder(place)
      });
      marker.addListener('click', function() {
        if(lastwindow)
          lastwindow.close();
        infowindow.open(map, marker);
        lastwindow = infowindow;
      });
      markers.push(marker);
    })});
    console.log("num windows = " + openwindows.length)
}
function sidepanel(mapsid){
  var request = {
    placeId: mapsid
  };
  service.getDetails(request, (r)=>{
      console.log(r.name);
      console.log(mapsid);
      $('.ui.sidebar').html("");
      $('.ui.sidebar').append("<h1>" + r.name + "</h1>");
      if(r.website)
        $('.ui.sidebar').append("<a href=\"" + r.website + "\">" + r.website + "</a>");
      if(r.international_phone_number)
        $('.ui.sidebar').append("<br>Phone: <a href=\"tel:" + r.international_phone_number + "\">" + r.international_phone_number + "</a>")
        $('.ui.sidebar').append("<h2>Reviews</h2>")
        $('.ui.sidebar').append("<div id=\"google-reviews\"></div>");
      $("#google-reviews").googlePlaces({placeId: mapsid, render: ['phone','reviews']});
      $('.ui.sidebar').sidebar('toggle');
  });
}

function infoWindowBuilder(place){
  var str = "";
  str += "<h1>" + place.barname + "</h1>";
  str += "<h2>" + place.etype + " - " + place.dayofweek + "</h2>";
  if(place.comments)
    str += "<p>" + place.comments + "</p>";
  str += "<a  href=\"#\" onclick=\"sidepanel(\`"+place.mapsid +"\`)\">More Info</a>";
  return str;
}
function getIcon(place){
  switch(place.etype){
    case "Trivia":
      return "trivia.png";
      break;
    case "Karaoke":
      return "karaoke.png";
      break;
    case "Open Mic":
      return "openmic.png";
      break;
    case "Comedy":
      return "comedy.png";
      break;
  }
}
