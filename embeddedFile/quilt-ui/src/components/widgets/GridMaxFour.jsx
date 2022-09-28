import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import { get, map } from 'lodash';

import GridMaxFourComponent from './GridMaxFourComponent'
import { parseQUIPrefix, getValueFromParsedObject, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    widgetContainer: {
        marginTop: "1rem",
        marginBottom: "1rem",
    },
    telcoDetailItemContainer: {
        margin: "auto",
        paddingRight: "1.5rem"
    },
    telcoDetailTitle: {
        ...typography.heading1,
        margin: "auto"
    },
    telcoDetailLabel: {
        ...typography.body1,
    },
    telcoDetailIcon: {
        maxWidth: "1.5rem",
        paddingRight: "0.25rem",
        float: "right"
    },
    telcoDetailTitleImage: {
        maxWidth: "1.1rem",
    },
    clickableCard: {
        cursor: "pointer",
        width: "100%",
        margin: "0.5rem",
        padding: "1rem",
        textAlign: "center"
    }
}));

const GridMaxFour = ({widget, updateEditorComponent, resetEditorComponent, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const [childComponents, setChildComponents] = useState([]);
    const [widgetData, setWidgetData] = useState({});
    const [editing, setEditing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [clickableCards, setClickableCards] = useState([]);
    // selectedCard state is used to track which AddonExpandableSection component is being selected
    const [selectedCard, setSelectedCard] = useState(null)
    const classes = useStyles();

    // If widget is selected for editing, render EditorContainer
    useEffect(() => {
        // Note that if a selectedCard is set, this GridMaxFour component will not update the editorComponent since the selected GridMaxFourComponent will be the one updating it instead.
        if (shouldRenderEditorComponent && !selectedCard && widgetData.length != 0) {
            // updateLocalisationValueFunc is a wrapper function for updateLocalisationValue to avoid having to pass in params(that are local to this scope) to renderEditorFieldComponents function.
            const updateLocalisationValueFunc = (localeKey, localeValue) => {
                updateLocalisationValue({localeKey, localeValue, pageId, setErrorMessage, setSaving, setEditing, fetchPageData})
            }
            // Generate EditorFieldComponents based on widgetData
            const editorFieldComponents = renderEditorFieldComponents({widgetData, saving, setWidgetData, editing, setEditing, initState, updateLocalisationValueFunc})
            // Invoke updateEditorComponent to render the EditorComponent
            updateEditorComponent(
                <EditorContainer 
                    errorMessage={errorMessage} 
                    editorFieldComponents={editorFieldComponents}
                    resetFunction={resetEditorComponent}
                    clickableCards={clickableCards}
                />
            )
        }
    }, [shouldRenderEditorComponent, widgetData, saving, errorMessage, selectedCard])

    // resetSelectedCard is passed as a prop to GridMaxFourComponent to unselect itself.
    const resetSelectedCard = () => {
        setSelectedCard(null)
    }

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    // It also instantiates the clickableCards state that is used to enable the user to select the card that they want to edit
    const initState = () => {
        if (!widget || get(widget, "components.length") === 0) return
        setChildComponents(null)
        setClickableCards([])
        setChildComponents(map(widget.components, component => {
            // Get cardTitle for displaying in ClickableCard
            let cardTitle = ""
            let uniqueKey = ""
            const componentType = get(component, "type")
            if (componentType === "telco-detail") {
                const {parsedKeyValue} = parseQUIPrefix(get(component, "data.image"))
                const imageUrl = getValueFromParsedObject(parsedKeyValue)
                cardTitle = <img src={imageUrl} alt="telco detail image"/>
                uniqueKey = imageUrl
            }
            const clickableCard = (<Card
                key={uniqueKey}
                className={clsx(classes.clickableCard)}
                onClick={()=>{
                    setSelectedCard(uniqueKey)
                }}
            >
                {cardTitle}
            </Card>)

            setClickableCards(prevState => [
                ...prevState,
                clickableCard
            ])

            const shouldRenderItemEditorComponent = selectedCard === uniqueKey
            return <GridMaxFourComponent
                key={uniqueKey}
                widget={component}
                updateEditorComponent={updateEditorComponent}
                shouldRenderEditorComponent={shouldRenderItemEditorComponent}
                resetSelectedCard={resetSelectedCard}
                pageId={pageId}
                fetchPageData={fetchPageData}
            />
        }))
    }

    // This effect is used to re-generate the clickableCards after a card has been clicked
    useEffect(() => {
        initState()
    }, [selectedCard])

    useEffect(() => {
        initState()
    }, [widget])

    if (get(widget, "components.length") ===0) {
        return <div>Empty {widget.type}</div>
    }

    return (<Box className={clsx(classes.widgetContainer)}>
        <Grid container direction="row">
            {map(childComponents, childComponent => childComponent)}
        </Grid>
    </Box>)
}

export default GridMaxFour;