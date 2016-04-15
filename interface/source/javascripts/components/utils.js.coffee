@Stats = React.createClass
  render: ->
    return (
      <div id="stats">
        
      </div>
    )

@Welcome = React.createClass
  render: ->
    <div className="welcome">
      <h1>OscFlow</h1>
      {@.props.children}
      <PatchesList/>
    </div>



@CheckBox = React.createClass

  render: ->
    return (

      <span className="mdl-list__item-secondary-action">
        <label className="demo-list-radio mdl-radio mdl-js-radio mdl-js-ripple-effect mdl-js-ripple-effect--ignore-events is-upgraded is-checked" 
          for="list-option-1" data-upgraded=",MaterialRadio,MaterialRipple">
          
          <input type="radio" 
            id="list-option-1" 
            className="mdl-radio__button" 
            name="options" 
            value="1" 
            checked=""
          />

          <span className="mdl-radio__outer-circle">
          </span>

          <span className="mdl-radio__inner-circle">
          </span>

          <span className="mdl-radio__ripple-container mdl-js-ripple-effect mdl-ripple--center"
           data-upgraded=",MaterialRipple">

           <span className="mdl-ripple is-animating"></span>
          </span>

        </label>
      </span>
    )