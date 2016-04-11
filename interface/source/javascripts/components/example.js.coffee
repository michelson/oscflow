@App = React.createClass

  render: ->
    return (
      <div className="mdl-layout mdl-js-layout mdl-layout--fixed-header">
        <header className="mdl-layout__header">
          <div className="mdl-layout__header-row">
            {# Title }
            <span className="mdl-layout-title">OscFlow</span>
            {# Add spacer, to align navigation to the right }
            <div className="mdl-layout-spacer"></div>
          </div>
        </header>

        <MainControl/>

        <main className="mdl-layout__content">
          <div className="page-content">

            <div id="stats">
              sin react 
            </div>

            {@props.children}

            <PatchesList/>
            <MainControl/>

          </div>
        </main>
      </div>
    )

@Welcome = React.createClass
  render: ->
    <div className="welcome">
      <h1>OscFlow</h1>
      {@.props.children}
      <PatchesList/>
    </div>

@PatchesList = React.createClass

  getInitialState: ->
    collection: []

  componentDidMount: ->

    $.getJSON( "http://localhost:8181/patches", ()=>
    ).done (data)=>
      @setState collection: data

  render: ->
    return(

      <section className="section--center mdl-grid mdl-grid--no-spacing mdl-shadow--2dp">
        <div className="mdl-card mdl-cell mdl-cell--12-col">
          
          <div className="mdl-card__supporting-text mdl-grid mdl-grid--no-spacing">
            
            <h4 className="mdl-cell mdl-cell--12-col">
              Patches list
            </h4>
            
            <div className="section__circle-container mdl-cell mdl-cell--2-col mdl-cell--1-col-phone">
              <div className="section__circle-container__circle mdl-color--primary"></div>
            </div>

            {
              @state.collection.map (item)=>

                <div className="section__text mdl-cell mdl-cell--10-col-desktop mdl-cell--6-col-tablet mdl-cell--3-col-phone">
                  <ReactRouter.Link to="/patches/1">
                    <h5>{item.name}</h5>
                    
                  </ReactRouter.Link>
                </div>

            }



          </div>

          <div className="mdl-card__actions">
            <a href="#" className="mdl-button">Read our features</a>
          </div>
        </div>
      </section>

    )

@MainControl = React.createClass
  render: ->
    <div className="mdl-layout__drawer">
      <span className="mdl-layout-title">OscFlow</span>
      <nav className="mdl-navigation">
        <a className="mdl-navigation__link" href="">Play</a>
        <a className="mdl-navigation__link" href="">Rec</a>
        <a className="mdl-navigation__link" href="">Status</a>
      </nav>
    </div>