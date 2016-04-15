@App = React.createClass

  render: ->
    
    return (
      <div className="mdl-layout mdl-js-layout mdl-layout--fixed-header">
        <Header/>

        <MainControl/>

        <main className="mdl-layout__content">
          <div className="page-content">

            {
              if @.props.location.pathname is "/"
                <PatchesList/>
            }

            {@props.children}
            
            <MainControl/>
            
            <Stats/>
          </div>
        </main>
      </div>
    )

