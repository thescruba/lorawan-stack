// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'

import PropTypes from 'prop-types'
import { Map, Marker, TileLayer } from 'react-leaflet'
import classnames from 'classnames'
import Leaflet from 'leaflet'
import bind from 'autobind-decorator'
import MarkerIcon from '../../assets/auxiliary-icons/location_pin.svg'

import style from './map.styl'

delete Leaflet.Icon.Default.prototype._getIconUrl
Leaflet.Icon.Default.mergeOptions({
  iconRetinaUrl: MarkerIcon,
  iconUrl: MarkerIcon,
  iconSize: [26, 36],
  shadowSize: [26, 36],
  iconAnchor: [13, 36],
  shadowAnchor: [8, 37],
  popupAnchor: [0, -40],
  shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
})

const defaultLocation = [52.3, 4.8]
const tileLayer = (
  <TileLayer
    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
    attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
  />
)
const defaultMinZoom = 7

export default class LocationMap extends React.Component {
  static propTypes = {
    className: PropTypes.string,
    // LeafletConfig is an object which can contain any number of properties defined by the leaflet plugin and is used to overwrite the default configuration of leaflet.
    leafletConfig: PropTypes.shape({}),
    // Markers is an array of objects containing a specific properties
    markers: PropTypes.arrayOf(
      // Position is a object containing two properties latitude and longitude which are both numbers.
      PropTypes.shape({
        position: PropTypes.shape({
          longitude: PropTypes.number,
          latitude: PropTypes.number,
        }),
      }),
    ),
    onClick: PropTypes.func,
    // Widget is a boolean used to add a class name to the map container div for styling.
    widget: PropTypes.bool,
  }

  static defaultProps = {
    leafletConfig: {},
    className: undefined,
    widget: false,
    markers: [],
    onClick: undefined,
  }

  constructor(props) {
    super(props)

    this.state = {
      zoomLevel: props.leafletConfig.zoom,
      mapCenter:
        props.markers.length >= 1
          ? [props.markers[0].position.latitude, props.markers[0].position.longitude]
          : this.getCurrentLocation(),
    }
  }

  getCurrentLocation() {
    if ('geolocation' in navigator) {
      navigator.geolocation.getCurrentPosition(function(position) {
        return [position.coords.latitude, position.coords.longitude]
      })
    }
    return defaultLocation
  }

  // fix the issue where tiles sometimes partially load
  componentDidMount() {
    const map = this.refs.map.leafletElement
    setTimeout(() => {
      map.invalidateSize()
    }, 250)
  }

  @bind
  handleClick(event) {
    const { onClick } = this.props
    const map = this.refs.map.leafletElement
    this.setState({ mapCenter: map.getCenter() })
    onClick(event)
  }

  @bind
  onZoomEvent(event) {
    this.setState({ zoomLevel: event.target._zoom })
  }

  render() {
    const { className, widget, markers, leafletConfig } = this.props
    const { zoomLevel, mapCenter } = this.state
    return (
      <div className={classnames(style.container, className)}>
        <Map
          ref="map"
          className={classnames(style.map, { [style.widget]: widget })}
          minZoom={defaultMinZoom}
          onZoomend={this.onZoomEvent}
          onClick={this.handleClick}
          {...leafletConfig}
          zoom={zoomLevel}
          center={mapCenter}
        >
          {tileLayer}
          {markers.map(
            marker =>
              marker && (
                <Marker
                  key={`marker-${marker.position.latitude}-${marker.position.altitude}`}
                  position={[marker.position.latitude, marker.position.longitude]}
                />
              ),
          )}
        </Map>
      </div>
    )
  }
}
