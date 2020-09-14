// Search Ingredients

function search() {
    const filter = $('#search').val().toUpperCase();
    $('.categories>li').each(function() {
        let show = false;
        $(this).children('ul').first().children('li').each(function() {
            const module = $(this).children('.name').first().text().toUpperCase();
            if(module.indexOf(filter) > -1) {
                show = true;
                $(this).css('display', 'block');
            } else {
                $(this).css('display', 'none');
            }
        });
        $(this).css('display', show ? 'block' : 'none');
    });
}

// Drag, Drop & Sortable Ingredients

$(".categories>li>ul").each(function() {
    Sortable.create(this, {
        group: {
        name: "ingredients",
        pull: "clone",
        },
        sort: false,
    });
});
  
Sortable.create($("#recipe")[0], {
    group: {
        name: "recipe",
        put: ["ingredients"],
    },
    animation: 150,
    chosenClass: "sorting",
    dragClass: "d-none",
    removeOnSpill: true,
});