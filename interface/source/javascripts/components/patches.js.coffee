@Patch = React.createClass

  getInitialState: ->
    resource: []

  componentDidMount: ->

    $.getJSON( "http://localhost:8181/patches/#{@props.params.patch}", ()=>
    ).done (data)=>
      console.log(data)
      @setState resource: data

  render: ->
    #console.log(@props.params.patch)
    return (
      <div>
        <h3>This the patch {@props.params.patch}</h3>
        <h3>This the patch {@state.resource.name}</h3>
      </div>
    ) 

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

            {
              @state.collection.map (item)=>
                <div className="section__circle-container mdl-cell mdl-cell--2-col mdl-cell--1-col-phone">
                  <div className="section__circle-container__circle mdl-color--primary"></div>
                </div>                
                <div className="section__text mdl-cell mdl-cell--10-col-desktop mdl-cell--6-col-tablet mdl-cell--3-col-phone">
                  <ReactRouter.Link to="/patches/#{item.name}">
                    <h5>{item.name}</h5>
                  </ReactRouter.Link>
                </div>
                
            }
          </div>

          <div className="mdl-card__actions">
            <a href="#" className="mdl-button">Sync Patches</a>
          </div>
        </div>
      </section>

    )

