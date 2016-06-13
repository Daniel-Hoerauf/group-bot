from flask import Flask, request, session, g, redirect, url_for, abort, render_template, flash
from pprint import PrettyPrinter

pp = PrettyPrinter(indent=2)

app = Flask(__name__)

@app.route('/callback_url')
def receive():
	pp.pprint(request)
	return "Hello World"
