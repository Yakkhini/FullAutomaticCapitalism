# frozen_string_literal: true

require 'binance'
require 'bundler/setup'

key = 'f63701669351442efdfeeae575ebb3dc8078838927631b5a7f0be4c0c1547302'
secret = 'a0a6f98ae868953d6eae5b04cbf65950bf0e55e738ed54fadf53cd37eadf6afe'

# Create a new client instance.
# If the APIs do not require the keys, (e.g. market data), key and secret can be omitted.
client = Binance::Spot.new(key: key,
                           secret: secret, base_url: 'https://testnet.binance.vision')

# Send a request to query server time
puts client.time

# Send a request to query BTCUSDT ticker
puts client.ticker_24hr(symbol: 'BTCUSDT')

# Send a request to get account information
puts client.account

# Place an order
response = client.new_order(symbol: 'BNBUSDT', side: 'BUY', price: 20, quantity: 1, type: 'LIMIT', timeInForce: 'GTC')
