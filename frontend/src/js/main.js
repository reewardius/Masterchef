// ==== Web Sockets ====

var wschannel = new WebSocket('ws://{{ .Host }}:{{ .Port }}/ws');

wschannel.onerror = function(error) {
    console.log('[-] Connection with the backend failed: ' + error);
};

wschannel.onmessage = function(message) {
    var data = message.data;
    var output = document.querySelector('#output');
    if(data === '%%#/clear/%%') {
        output.value = '';
    } else {
        output.value = data;
    };
}

// ==== Bake ====

document.querySelector('#bake').addEventListener('click', function(e) {
    e.preventDefault();
    var cooker = { input: '', recipe: [] };
    // Check the target
    var input = document.querySelector('#input').value.trim();
    cooker['input'] = input;
    // Get the recipe    
    document.querySelectorAll('#recipe > li').forEach(function(el) {
        // Check state
        if(el.getAttribute('off') === 'true') {
            return;
        }
        // Get the ingredient
        var ingredient = { name: '', incognito: false, single: true, calories: {} };
        ingredient['name'] = el.querySelector('.name').innerText;
        // Check single
        ingredient['single'] = el.classList.contains('single');
        // Check incognito
        ingredient['incognito'] = el.getAttribute('incognito') === 'true';
        // Get the calories
        if(!ingredient['incognito']) {
            el.querySelector('.calories').childNodes.forEach(function(child) {
                var name = child.querySelector('label').innerText;
                var value = child.querySelector('input').value;
                ingredient['calories'][name] = value;
            });
        }
        cooker['recipe'].push(ingredient);
    });
    // Send to the web socket
    var message = '#/cook/' + JSON.stringify(cooker);
    wschannel.send(message);
});

// ==== Options ====

function ingredient2recipe(el) {
    var options = document.createElement('div');
    options.classList.add('opts');
    // Icognito mode (using OSINT)
    if(el.classList.contains('osint')) {
        el.classList.remove('osint');
        el.parentNode.setAttribute('incognito', 'false');
        var icognito = document.createElement('i');
        icognito.classList.add('icon-incognito');
        icognito.onclick = function(e) {
            e.preventDefault();
            var ell = e.target.parentNode.parentNode.parentNode;
            var turnOn = ell.getAttribute('incognito') === 'false';
            ell.setAttribute('incognito', turnOn ? 'true' : 'false');
        }
        options.append(icognito);
    }
    // Disable button
    el.parentNode.setAttribute('off', 'false');
    var disabled = document.createElement('i');
    disabled.classList.add('icon-disable');
    disabled.onclick = function(e) {
        e.preventDefault();
        var ell = e.target.parentNode.parentNode.parentNode;
        var turnOn = ell.getAttribute('off') === 'false';
        ell.setAttribute('off', turnOn ? 'true' : 'false');
    }
    options.append(disabled);
    // Create options
    el.append(options);
}

// ==== Search Ingredients ====

function search() {
    var filter = document.querySelector('#search').value.toUpperCase();
    document.querySelector('#search-results').innerHTML = '';
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

document.querySelectorAll('.categories>li>ul').forEach(function(el) {
    Sortable.create(el, {
        group: {
            name: 'ingredients',
            put: false,
            pull: 'clone',
        },
        sort: false,
    });
});
Sortable.create(document.getElementById('search-results'), {
    group: {
        name: 'ingredients',
        put: false,
        pull: 'clone',
    },
    sort: false,
});

Sortable.create(document.getElementById('recipe'), {
    group: {
        name: 'recipe',
        put: ['ingredients'],
    },
    animation: 150,
    sort: true,
    removeOnSpill: true,
    onAdd: function(evt) { ingredient2recipe(evt.item.querySelector('.name')) },
});