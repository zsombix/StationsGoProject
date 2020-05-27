const stationList = document.querySelector('.collection');

var stations;


//Load all event listeners
loadEventListiners();

//Init the call to the API
loadStations()

function loadEventListiners(){

}

function loadStations(){
  fetch('http://localhost:8080/stations')
  .then(response => response.json())
  .then(data => {		
		stations = data;
		stations.forEach(station => {
			//Create list element
			const li = document.createElement('li');
			li.className = 'collection-item';
			//Add the text
			li.appendChild(document.createTextNode(station.Name));
			//Add to the List
			stationList.appendChild(li);
		});

		
	});
}