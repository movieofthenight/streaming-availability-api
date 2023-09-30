let chosenLogos = document.getElementById("chosen-logos");
let notChosenLogos = document.getElementById("not-chosen-logos");
let mainPage = document.getElementById("main-page");
let serviceSelection = document.getElementById("service-selection");
let serviceSelectionContainer = document.getElementById("service-selection-container")
let saveButton = document.getElementById("save-button");
let form = document.getElementById("options-form");
let genreInput = document.getElementById("genre-input");
let movieTypeInput = document.getElementById("movie-type-input");
let keywordInput = document.getElementById("keyword-input");
let countryInput = document.getElementById("country-input");
let results = document.getElementById("results");

saveButton.addEventListener("click", () => {
	serviceSelectionContainer.style.display = "none";
	mainPage.style.display = "";
});

form.addEventListener("submit", (e) => {
	e.preventDefault();
	let genre = genreInput.value;
	let movieType = movieTypeInput.value;
	let services = Service.enabledServices;
	let country = countryInput.value;
	let keyword = keywordInput.value
	fetch(`api/recommend?genre=${genre}&movieType=${movieType}&keyword=${encodeURIComponent(keyword)}&services=${services.join(",")}&country=${country}`).then((res) => {
		return res.json()
	}).then((res) => {
		results.innerHTML = "";
		res.forEach((movie) => {
			let movieDiv = document.createElement("a");
			movieDiv.href = movie.streamingLink;
			movieDiv.target = "_blank";
			movieDiv.classList.add("movie");
			let movieImg = document.createElement("img");
			movieImg.src = movie.poster;
			movieImg.alt = movie.title;
			movieImg.classList.add("movie-poster");
			movieDiv.appendChild(movieImg);
			let streamingLogo = document.createElement("img");
			streamingLogo.src = movie.streamingLogo;
			streamingLogo.classList.add("streaming-poster-logo");
			movieDiv.appendChild(streamingLogo);
			results.appendChild(movieDiv);
		});
	});
});

document.getElementById("services-option").addEventListener("click", () => {
	mainPage.style.display = "none";
	serviceSelectionContainer.style.display = "";
});


class Service {

	static enabledServices = Service.getLocalServices();

	static getLocalServices() {
		let localServices = localStorage.getItem("services");
		if(localServices == null) {
			// Default services if none are set
			return ["netflix", "disney", "max", "hulu", "prime", "hbo"];
		}
		return JSON.parse(localServices);
	}

	constructor(id, name, darkThemeLogo, whiteLogo, themeColor) {
		this.id = id;
		this.name = name;
		this.darkThemeLogo = darkThemeLogo;
		this.whiteLogo = whiteLogo;
		this.themeColor = themeColor;

		this.homePageChosenImg = document.createElement("img");
		this.homePageChosenImg.src = this.darkThemeLogo;
		this.homePageChosenImg.alt = this.name;
		this.homePageChosenImg.style.display = "none";
		this.homePageChosenImg.classList.add("service-logo");
		chosenLogos.appendChild(this.homePageChosenImg);

		this.homePageNotChosenImg = document.createElement("img");
		this.homePageNotChosenImg.src = this.darkThemeLogo;
		this.homePageNotChosenImg.alt = this.name;
		this.homePageNotChosenImg.style.display = "none";
		this.homePageNotChosenImg.classList.add("service-logo");
		notChosenLogos.appendChild(this.homePageNotChosenImg);

		this.serviceSelectionChild = document.createElement("div");
		this.serviceSelectionChild.classList.add("service-selection-child");
		this.serviceSelectionChild.addEventListener("click", () => {
			if(this.enabled) {
				this.disable();
			} else {
				this.enable();
			}
		});

		this.serviceSelectionImg = document.createElement("img");
		this.serviceSelectionImg.src = this.whiteLogo;
		this.serviceSelectionImg.alt = this.name;
		this.serviceSelectionImg.classList.add("service-list-logo");

		this.nameDiv = document.createElement("div");
		this.nameDiv.textContent = this.name;
		this.nameDiv.classList.add("service-name");

		this.addOrDoneImg = document.createElement("img");
		this.addOrDoneImg.classList.add("add-or-done");

		this.serviceSelectionChild.append(this.serviceSelectionImg, this.nameDiv, this.addOrDoneImg);
		serviceSelection.appendChild(this.serviceSelectionChild);

		if(Service.enabledServices.includes(this.id)) {
			this.enable();
		} else {
			this.disable();
		}
	}

	enable() {
		this.homePageChosenImg.style.display = "block";
		this.homePageNotChosenImg.style.display = "none";
		this.addOrDoneImg.src = "static/done.svg";
		this.addOrDoneImg.style.filter = "";
		this.serviceSelectionImg.style.filter = "";
		this.serviceSelectionChild.style.backgroundColor = this.themeColor;
		this.serviceSelectionChild.style.color = "var(--theme-color-7)";
		this.enabled = true;
		if(!Service.enabledServices.includes(this.id)) {
			Service.enabledServices.push(this.id);
		}
		setTimeout(() => {
			localStorage.setItem("services", JSON.stringify(Service.enabledServices));
		}, 1);
	}

	disable() {
		this.homePageChosenImg.style.display = "none";
		this.homePageNotChosenImg.style.display = "block";
		this.addOrDoneImg.src = "static/add.svg";
		this.addOrDoneImg.style.filter = "contrast(0)";
		this.serviceSelectionImg.style.filter = "contrast(0)";
		this.serviceSelectionChild.style.backgroundColor = "";
		this.serviceSelectionChild.style.color = "var(--theme-color-4)";
		this.enabled = false;
		if(Service.enabledServices.includes(this.id)) {
			Service.enabledServices.splice(Service.enabledServices.indexOf(this.id), 1);
		}
		setTimeout(() => {
			localStorage.setItem("services", JSON.stringify(Service.enabledServices));
		}, 1);
	}
}

fetch("api/countries").then((res) => {
	return res.json();
}).then((res) => {
	let countryOptions = [];
	for(let countryCode in res) {
		let country = res[countryCode];
		let option = document.createElement("option");
		option.value = countryCode;
		option.innerText = country.name;
		countryOptions.push(option);
	}
	countryOptions.sort((a, b) => {
		return a.innerText.localeCompare(b.innerText);
	});
	countryInput.addEventListener("change", () => {
		chosenLogos.innerHTML = "";
		notChosenLogos.innerHTML = "";
		serviceSelection.innerHTML = "";
		let services = [];
		for(let serviceId in res[countryInput.value].services) {
			services.push(res[countryInput.value].services[serviceId]);
		}
		services.sort((a, b) => {
			return a.name.localeCompare(b.name);
		});
		services.forEach((service) => {
			new Service(service.id, service.name, service.darkThemeLogo, service.whiteLogo, service.themeColor);
		});
		setTimeout(() => {
			localStorage.setItem("country", countryInput.value);
		}, 1);
	});
	countryInput.append(...countryOptions);
	getInitialCountry.then((country) => {
		if(res[country] != null) {
			return country;
		} else {
			return Promise.reject()
		}
	}).catch(() => {
		return "us"
	}).then((country) => {
		countryInput.value = country;
		countryInput.dispatchEvent(new Event('change'));
		setTimeout(() => {
			localStorage.setItem("country", country);
		}, 1);
	});
});

fetch("api/genres").then((res) => {
	return res.json();
}).then((res) => {
	let genreOptions = [];
	res.forEach((genre) => {
		let option = document.createElement("option");
		option.value = genre.id;
		option.innerText = genre.name;
		genreOptions.push(option);
	});
	genreOptions.sort((a, b) => {
		return a.innerText.localeCompare(b.innerText);
	});
	genreInput.append(...genreOptions);
});

fetch("api/movie-types").then((res) => {
	return res.json();
}).then((res) => {
	let movieTypeOptions = [];
	res.forEach((movieType) => {
		let option = document.createElement("option");
		option.value = movieType.id;
		option.innerText = movieType.name;
		movieTypeOptions.push(option);
	});
	movieTypeInput.append(...movieTypeOptions);
});

let getInitialCountry = new Promise((resolve, reject) => {
	let country = localStorage.getItem("country");
	if(country != null) {
		resolve(country);
		return;
	}
	fetch("api/user-country").then((res) => {
		return res.json();
	}).then((res) => {
		if(res.detected) {
			resolve(res.country);
		} else {
			reject();
		}
	}).catch(() => {
		reject();
	});
});
