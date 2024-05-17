from flask import Flask, request, jsonify, abort

app = Flask(__name__)

orders = {}
next_id = 1

@app.route('/novo', methods=['POST'])
def create_order():
    global next_id
    data = request.get_json()
    
    id = next_id
    orders[id] = {
        'id': id,
        'nome': data['nome'],
        'email': data['email'],
        'descricao': data['descricao']
    }

    next_id += 1

    return jsonify({'id': id}), 200

@app.route('/pedidos', methods=['GET'])
def get_orders():
    return jsonify(list(orders.values()))

@app.route('/pedidos/<int:id>', methods=['GET'])
def get_order_by_id(id):
    order = orders.get(id)
    if order is None:
        abort(404, 'Order not found')
    return jsonify(order)

@app.route('/pedidos/<int:id>', methods=['PUT'])
def update_order(id):
    order = orders.get(id)
    if order is None:
        abort(404, 'O pedido não existe')

    data = request.get_json()

    order.update({
        'nome': data['nome'],
        'email': data['email'],
        'descricao': data['descricao']
    })
    return jsonify(order)

@app.route('/pedidos/<int:id>', methods=['DELETE'])
def delete_order(id):
    if id not in orders:
        abort(404, 'O pedido não existe')
    
    del orders[id]
    return jsonify('Pedido excluído com sucesso!'), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
