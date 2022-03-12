




let search_input = document.getElementById('search_input');
let liElements = document.getElementsByClassName('option')
let suggestions_block = document.getElementsByClassName('suggestions_block')[0]

search_input.addEventListener(
    'input',
    show_suggestion
);

window.addEventListener(
    'click',
    function bodyclick(e) {
        suggestions_block.style.cssText = "display: none;"
    }
)

for (let i = 0; i < liElements.length; i++) {
    liElements[i].addEventListener(
        'click',
        selectOption
    )
};


function selectOption(e) { 
    search_input.focus()
    let search_text = e.target.innerText.split(" // ")[0]
    let search_type =  e.target.innerText.split(" // ")[1]
    let select =  document.getElementById('select')
    select.value = search_type
    search_input.value = search_text
    console.log(search_text,search_type)
};

function show_suggestion(e) {
    let liElements = document.getElementsByClassName('option')
    let search_text = e.target.value.toLowerCase()
    let arrayofSuggestion = []
    for (let i = 0; i < liElements.length; i++) {
        let elemText = liElements[i].innerHTML.toLowerCase()
        if (elemText.includes(search_text)) {
            let el = liElements[i]
            arrayofSuggestion.push(el)
            liElements[i].style.cssText = "display: block;"
        } else {
            liElements[i].style.cssText = "display: none;"
        }
            
    }

    suggestions_block.style.cssText = "display: block;"
};


function showfilter() {
    filter_form = document.getElementById("filter_form")
    checker =  document.getElementById("checker")
    if (filter_form.style.display == 'none' || filter_form.style.display == "" )  {
        filter_form.style.display = 'block'
        checker.checked = true

    } else {
        filter_form.style.display = 'none'
        checker.checked = false

    }
} 

let filtersync = document.getElementById('filtersync')

filtersync.addEventListener(
    'input',
    syncrange
);

let rangeInput = document.getElementById('rangeInput')
rangeInput.disabled = true;
function syncrange(e) {
    let val = filtersync.value
    let rangeInput = document.getElementById('rangeInput')
    rangeInput.value  = val
    console.log(val)
    
}