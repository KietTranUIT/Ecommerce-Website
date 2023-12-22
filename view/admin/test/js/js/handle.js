$(document).ready(function() {
    $('.category-btn').click(function() {
        console.log('ok')
        var id = this.dataset.categoryId

        var url = '/categories/' + id
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
            window.location.href = response.href
        })
    })
})
