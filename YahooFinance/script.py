import numpy as np
import pandas as pd
import yfinance as yf
import plotly.graph_objs as go

# Interval required 5 minutes
data = yf.download(tickers='GME', period='5d', interval='5m')
# Print data
data.to_csv('GME.csv')

# creata table in html
a = pd.read_csv("GME.csv")
a.to_html("table.html")

# create graph
fig = go.Figure()

# Candlestick
fig.add_trace(go.Candlestick(x=data.index,
                             open=data['Open'],
                             high=data['High'],
                             low=data['Low'],
                             close=data['Close'], name='market data'))

# Add titles
fig.update_layout(
    title='GME stock',
    yaxis_title='Stock Price (USD per Shares)')

# X-Axes
fig.update_xaxes(
    rangeslider_visible=True,
    rangeselector=dict(
        buttons=list([
            dict(count=15, label="15m", step="minute", stepmode="backward"),
            dict(count=45, label="45m", step="minute", stepmode="backward"),
            dict(count=1, label="HTD", step="hour", stepmode="todate"),
            dict(count=3, label="3h", step="hour", stepmode="backward"),
            dict(step="all")
        ])
    )
)

# Show
fig.show()
