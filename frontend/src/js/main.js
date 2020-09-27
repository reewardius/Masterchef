// ==== Bake ====

var wschannel = new WebSocket("ws://{{ .Host }}:{{ .Port }}/ws");

wschannel.onerror = function(error) {
    console.log("[-] Connection with the backend failed: " + error);
};

wschannel.onmessage = function(message) {
    var data = message.data;
    var output = document.querySelector('#output');
    if(data === '%%#/clear/%%') {
        output.value = "";
    } else {
        output.value = data;
    };
}

// ==== Bake ====


document.querySelector('#bake').addEventListener('click', function(e) {
    e.preventDefault();
    var cooker = { input: "", recipe: [] };
    // Check the target
    var input = document.querySelector('#input').value.trim();
    cooker['input'] = input;
    // Get the recipe    
    document.querySelectorAll('#recipe > li').forEach(function(el) {
        // Get the ingredient
        var ingredient = { name: "", calories: [] };
        ingredient['name'] = el.querySelector('.name').innerText;
        // Get the calories
        var calories = {};
        el.querySelector('.calories').childNodes.forEach(function(child) {
            var name = child.querySelector('label').innerText;
            var value = child.querySelector('input').value;
            calories[name] = value;
        });
        ingredient['calories'] = calories;
        cooker['recipe'].push(ingredient);
    });
    // Send to the web socket
    var message = '#/cook/' + JSON.stringify(cooker);
    wschannel.send(message);
});

// ==== Options ====

document.querySelectorAll('.categories .name').forEach(function(el) {
    var options = document.createElement('div');
    options.classList.add('opts');
    // Stealth modality (using OSINT)
    if(el.classList.contains('osint')) {
        el.classList.remove('osint');
        var osint = document.createElement('i');
        osint.classList.add('fas');
        osint.classList.add('fa-user-secret');
        osint.setAttribute('on', 'false');
        osint.addEventListener('click', function(e) {
            e.preventDefault();
            var ell = e.target;
            var off = ell.getAttribute === "false";
            ell.setAttribute(off ? "true" : "false");
            ell.style = 'color: ' + (off ? "#468847" : "#aaa");
        });
        options.append(osint);
    }
    // Disable button
    var disabled = document.createElement('i');
    disabled.classList.add('fas');
    disabled.classList.add('fa-user-secret');
    disabled.setAttribute('on', 'false');
    disabled.addEventListener('click', function(e) {
        e.preventDefault();
        var ell = e.target;
        var off = ell.getAttribute === "false";
        ell.setAttribute(off ? "true" : "false");
        ell.style = 'color: ' + (off ? "#aaa" : "#ff0000");
        var name = ell.parentNode.parentNode;
        name.style = 'color' +  (off ? '#468847' : '#aaa');
        name.parentNode.style = 'background-color' + (off ? '#dff0d8': '#ddd');
    });
    options.append(disabled);
    // Create options
    el.append(options);
});

// ==== Search Ingredients ====

function search() {
    var filter = document.querySelector('#search').value.toUpperCase();
    document.querySelector('#search-results').innerHTML = "";
    if(filter.length > 0) {
        document.querySelectorAll('.categories>li>ul>li').forEach(function(module) {
            if(module.innerText.toUpperCase().indexOf(filter) > -1) {
                var cl = module.cloneNode(10);
                document.querySelector('#search-results').append(cl);
            }
        });
    }
}

// ==== Drag, Drop & Sortable Ingredients ====

document.querySelectorAll(".categories>li>ul").forEach(function(el) {
    Sortable.create(el, {
        group: {
            name: "ingredients",
            put: false,
            pull: "clone",
        },
        sort: false,
    });
});
Sortable.create(document.getElementById("search-results"), {
    group: {
        name: "ingredients",
        put: false,
        pull: "clone",
    },
    sort: false,
});

Sortable.create(document.getElementById("recipe"), {
    group: {
        name: "recipe",
        put: ["ingredients"],
    },
    animation: 150,
    sort: true,
    removeOnSpill: true,
});