// Options

$('.categories .name').each(function() {
    let options = $('<div class="opts"></div>');
    // Stealth modality (using OSINT)
    if($(this).hasClass('osint')) {
        $(this).toggleClass('osint');
        let osint = $('<i class="fas fa-user-secret" on="false">');
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
    let disabled = $('<i class="fas fa-ban" on="false">');
    disabled.on('click', function(e) {
        e.preventDefault();
        $(this).attr('on', function(index, attr) {
            return attr == 'false' ? 'true' : 'false';
        });
        const on = $(this).attr('on') == 'true';
        $(this).css('color', on ? '#ff0000' : '#aaa');
        const name = $(this).parent().parent();
        name.css('color', on ? '#aaa' : '#468847');
        name.parent().css('background-color', on ? '#ddd' : '#dff0d8');
    });
    options.append(disabled);
    // Create options
    $(this).append(options);
});

// Search Ingredients

function search() {
    const filter = $('#search').val().toUpperCase();
    $('#search-results').empty();
    if(filter.length > 0) {
        console.log("Filter length", filter);
        $('.categories>li>ul>li').each(function() {
            const module = $(this);
            if(module.text().toUpperCase().indexOf(filter) > -1) {
                $(this).clone().appendTo('#search-results');
            }
        });
    }
}

// Drag, Drop & Sortable Ingredients

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