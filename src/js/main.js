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