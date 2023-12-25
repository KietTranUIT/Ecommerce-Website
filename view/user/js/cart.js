var cartId = "cart";

var localAdapter = {

    saveCart: function (object) {

        var stringified = JSON.stringify(object);
        localStorage.setItem(cartId, stringified);
        return true;

    },
    getCart: function () {

        return JSON.parse(localStorage.getItem(cartId));

    },
    clearCart: function () {

        localStorage.removeItem(cartId);

    }

};

let storage = localAdapter


var cart = {

    count: 0,
    total: 0,
    items: [],
    getItems: function () {

        return this.items;

    },
    setItems: function (cart_id) {

        this.items = cart_id.items;
        this.count = cart_id.count;
        this.total = cart_id.total;
        console.log(cart)
    },
    // clearItems: function () {

    //     this.items = [];
    //     this.total = 0;
    //     storage.clearCart();
    //     helpers.emptyView();

    // },
    addItem: function (item) {

        if (this.containsItem(item.id) === false) {

            this.items.push({
                id: item.id,
                name: item.name,
                price: item.price,
                image: item.image,
                size: item.size,
                quantity: item.quantity,
                total: item.total
            });
            this.total += item.total
            this.count += item.quantity

            storage.saveCart(this);


        } else {

            this.updateItem(item);

        }
        $('#cart-count').text(`Cart(${this.count})`)

    },
    containsItem: function (id) {

        if (this.items === undefined) {
            return false;
        }

        for (var i = 0; i < this.items.length; i++) {

            var _item = this.items[i];

            if (id == _item.id) {
                return true;
            }

        }
        return false;

    },
    updateItem: function (object) {

        for (var i = 0; i < this.items.length; i++) {

            var _item = this.items[i];

            if (object.id == _item.id) {

                _item.quantity = parseInt(object.quantity) + parseInt(_item.quantity);
                _item.total = parseInt(object.total) + parseInt(_item.total);
                this.items[i] = _item;
                this.count += parseInt(object.quantity)
                this.total += parseInt(object.total)
                storage.saveCart(this);
                $('#cart-count').text(`Cart(${this.count})`)
            }

        }

    }

};


$(document).ready(function() {
    //let cart = localStorage.getItem('cart')

    if (storage.getCart()) {
        cart.setItems(storage.getCart())

        $('#cart-count').text(`Cart(${cart.count})`)
    } else {
        $('#cart-count').text('Cart(0)')
    }


    $('#cart-click').click(function() {
        let size_product = $("input[name='size']:checked");
        if (size_product.length == 0) {
            alert("Please choose a size product")
            return
        }

        let item = {
            id: parseInt(size_product.val()),
            size: parseInt(size_product.data('size')),
            image: $(this).data('product-image'),
            name: $(this).data('product-name'),
            price: $(this).data('product-price'),
            quantity: parseInt($("input[name='quantity']").val()),
            total: parseInt($(this).data('product-price')) * parseInt($("input[name='quantity']").val())
        }

        cart.addItem(item)
        alert("Added Item to Cart!")
    })

    $('.price-format').each(function() {
        let price = parseInt($(this).text())
        $(this).text(price.toLocaleString('it-IT', {style : 'currency', currency : 'VND'}))
    })

    
})

