// function place (id, mapsID, Barname, etype, dayofweek, lat, long){
//   this.id = id;
//   this.mapsID = mapsID;
//   this.Barname = Barname;
//   this.etype = etyps;
//   this.dayofweek = dayofweek;
//   this.lat = lat;
//   this.long = long;
// }
var deletedPlaces = [];

window.onload = ()=>{
  var id = 1;
  $("#button").click(()=>{
    if(deletedPlaces.length == 0) return;
    var row =  deletedPlaces.pop();
    $.ajax({
        url: '/places',
        type: 'PUT',
        data: JSON.stringify(row),
        success: function(result) {
            //console.log(result);
        }
    });
    $(".results").tabulator("setData", "../places");
  });

  $(".results").tabulator({
    height:205, // set height of table, this enables the Virtual DOM and improves render speed dramatically (can be any valid css height value)
    layout:"fitColumns", //fit columns to width of table (optional)
    history:true,
    columns:[ //Define Table Columns
        {title:"ID", field:"id", width:150},
        {title:"MapsID", field:"mapsid", width:150},
        {title:"BarName", field:"barname"},
        {title:"Event Type", field:"etype"},
        {title:"Day of Week", field:"dayofweek"},
        {title:"Lat", field:"lat"},
        {title:"long", field:"long"},
        {title:"Comments", field:"comments"},
        {title:"Delete", cellClick:function(e, cell){
          deleteRow(cell.getRow());
        }}
    ],
    // rowClick:function(e, row){ //trigger an alert message when the row is clicked
    //     console.log(row.getData());
    //     row.delete();
    // },
  });
  $(".results").tabulator("setData", "../places");

}
function deleteRow(row){
    $.ajax({
        url: '../places/' + row.getData().id,
        type: 'DELETE',
        success: function(result) {
            if(result == 0) alert("nothing to delete");
            else{
              deletedPlaces.push(row.getData());
              row.delete();
            }
        }
    });
}
