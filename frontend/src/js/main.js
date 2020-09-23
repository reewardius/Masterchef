// Dropzone
$('.filehere').each(function() {
    $(this).dropzone({ url: '#' });
});

// Options

$('.categories .name').each(function() {
    let options = $('<div class="opts"></div>');
    // Stealth modality (using OSINT)
    if($(this).hasClass('osint')) {
        $(this).toggleClass('osint');
        let osint = $('<i on="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M383.9 308.3l23.9-62.6c4-10.5-3.7-21.7-15-21.7h-58.5c11-18.9 17.8-40.6 17.8-64v-.3c39.2-7.8 64-19.1 64-31.7 0-13.3-27.3-25.1-70.1-33-9.2-32.8-27-65.8-40.6-82.8-9.5-11.9-25.9-15.6-39.5-8.8l-27.6 13.8c-9 4.5-19.6 4.5-28.6 0L182.1 3.4c-13.6-6.8-30-3.1-39.5 8.8-13.5 17-31.4 50-40.6 82.8-42.7 7.9-70 19.7-70 33 0 12.6 24.8 23.9 64 31.7v.3c0 23.4 6.8 45.1 17.8 64H56.3c-11.5 0-19.2 11.7-14.7 22.3l25.8 60.2C27.3 329.8 0 372.7 0 422.4v44.8C0 491.9 20.1 512 44.8 512h358.4c24.7 0 44.8-20.1 44.8-44.8v-44.8c0-48.4-25.8-90.4-64.1-114.1zM176 480l-41.6-192 49.6 32 24 40-32 120zm96 0l-32-120 24-40 49.6-32L272 480zm41.7-298.5c-3.9 11.9-7 24.6-16.5 33.4-10.1 9.3-48 22.4-64-25-2.8-8.4-15.4-8.4-18.3 0-17 50.2-56 32.4-64 25-9.5-8.8-12.7-21.5-16.5-33.4-.8-2.5-6.3-5.7-6.3-5.8v-10.8c28.3 3.6 61 5.8 96 5.8s67.7-2.1 96-5.8v10.8c-.1.1-5.6 3.2-6.4 5.8z"/></svg></i>');
        osint.on('click', function(e) {
            e.preventDefault();
            $(this).attr('on', function(index, attr) {
                return attr == "false" ? "true" : "false";
            });
            $(this).css('fill', $(this).attr('on') == "true" ? "#468847" : "#aaa");
        });
        options.append(osint);
    }
    // Disable button
    let disabled = $('<i on="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M256 8C119.034 8 8 119.033 8 256s111.034 248 248 248 248-111.034 248-248S392.967 8 256 8zm130.108 117.892c65.448 65.448 70 165.481 20.677 235.637L150.47 105.216c70.204-49.356 170.226-44.735 235.638 20.676zM125.892 386.108c-65.448-65.448-70-165.481-20.677-235.637L361.53 406.784c-70.203 49.356-170.226 44.736-235.638-20.676z"/></svg></i>');
    disabled.on('click', function(e) {
        e.preventDefault();
        $(this).attr('on', function(index, attr) {
            return attr == 'false' ? 'true' : 'false';
        });
        const on = $(this).attr('on') == 'true';
        $(this).css('fill', on ? '#ff0000' : '#aaa');
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