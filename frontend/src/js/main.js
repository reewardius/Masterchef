// ==== Bake ====

$('#bake').on('click', function() {
    var cooker = { input: "", recipe: [] };
    // Check the target
    var input = $('#input').val().trim();
    cooker['input'] = input;
    // Get the recipe    
    $('#recipe > li').each(function() {
        // Get the ingredient
        var ingredient = { name: "", calories: [] };
        ingredient['name'] = $(this).find('.name')[0].innerText;
        // Get the calories
        $(this).find('.calories').first().children().each(function() {
            var calories = { name: "", value: "" };
            calories['name'] = $(this).find('label')[0].innerText;
            calories['value'] = $(this).find('input').first().val();
            ingredient['calories'].push(calories);
        });
        cooker['recipe'].push(ingredient);
    });
    // Send to the web socket
    var _ = JSON.stringify(cooker);
    console.log(_);
});

// ==== Options ====

$('.categories .name').each(function() {
    var options = $('<div class="opts"></div>');
    // Stealth modality (using OSINT)
    if($(this).hasClass('osint')) {
        $(this).toggleClass('osint');
        var osint = $('<i class="fas fa-user-secret" on="false">');
        osint.on('click', function(e) {
            e.preventDefault();
            $(this).attr('on', function(index, attr) {
                return attr == "false" ? "true" : "false";
            });
            $(this).css('color', $(this).attr('on') == "true" ? "#468847" : "#aaa");
        });
        options.append(osint);
    }
    // Disable button
    var disabled = $('<i class="fas fa-ban" on="false">');
    disabled.on('click', function(e) {
        e.preventDefault();
        $(this).attr('on', function(index, attr) {
            return attr == 'false' ? 'true' : 'false';
        });
        var on = $(this).attr('on') == 'true';
        $(this).css('color', on ? '#ff0000' : '#aaa');
        var name = $(this).parent().parent();
        name.css('color', on ? '#aaa' : '#468847');
        name.parent().css('background-color', on ? '#ddd' : '#dff0d8');
    });
    options.append(disabled);
    // Create options
    $(this).append(options);
});

// ==== Search Ingredients ====

function search() {
    var filter = $('#search').val().toUpperCase();
    $('#search-results').empty();
    if(filter.length > 0) {
        console.log("Filter length", filter);
        $('.categories>li>ul>li').each(function() {
            var module = $(this);
            if(module.text().toUpperCase().indexOf(filter) > -1) {
                $(this).clone().appendTo('#search-results');
            }
        });
    }
}

// ==== Drag, Drop & Sortable Ingredients ====

$(".categories>li>ul").each(function() {
    Sortable.create(this, {
        group: {
            name: "ingredients",
            put: false,
            pull: "clone",
        },
        sort: false,
    });
});
Sortable.create($("#search-results")[0], {
    group: {
        name: "ingredients",
        put: false,
        pull: "clone",
    },
    sort: false,
});

Sortable.create($("#recipe")[0], {
    group: {
        name: "recipe",
        put: ["ingredients"],
    },
    animation: 150,
    sort: true,
    removeOnSpill: true,
});