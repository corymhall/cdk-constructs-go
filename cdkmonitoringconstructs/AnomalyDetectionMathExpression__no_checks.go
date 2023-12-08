//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AnomalyDetectionMathExpression) validateCreateAlarmParameters(scope constructs.Construct, id *string, props *awscloudwatch.CreateAlarmOptions) error {
	return nil
}

func (a *jsiiProxy_AnomalyDetectionMathExpression) validateWithParameters(props *awscloudwatch.MathExpressionOptions) error {
	return nil
}

func validateNewAnomalyDetectionMathExpressionParameters(props *awscloudwatch.MathExpressionProps) error {
	return nil
}

